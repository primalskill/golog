package golog

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"runtime"
	"sync"
	"time"
)

var (
	colorBgGreen = []byte("\x1b[42m")
	colorFgGreen = []byte("\x1b[32m")
	colorBgRed   = []byte("\x1b[41m")
	colorFgRed   = []byte("\x1b[31m")
	colorBgCyan  = []byte("\x1b[46m")
	colorFgCyan  = []byte("\x1b[36m")
	colorFgBlack = []byte("\x1b[30m")
	colorReset   = []byte("\x1b[0m")
	colorFaint   = []byte("\x1b[2m")
)

type groupOrAttrs struct {
	group string
	attrs []slog.Attr
}

type DevHandler struct {
	goas []groupOrAttrs
	mu   *sync.Mutex
}

func newDevHandler() *DevHandler {
	dh := &DevHandler{
		mu: &sync.Mutex{},
	}

	return dh
}

func (p *DevHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= slog.LevelDebug
}

func (p *DevHandler) Handle(_ context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)

	// append the level and the message
	buf, _, fgColor := p.appendLevelMessage(buf, r.Level, r.Message)

	// append time
	if r.Time.IsZero() != true {
		buf = p.appendTime(buf, r.Time, fgColor)
	}

	// append the source (file:line)
	if r.PC != 0 {
		buf = p.appendSource(buf, r.PC, fgColor)
	}

	buf = append(buf, '\n')

	// parse WithAttr and WithGroup
	goas := p.goas

	if r.NumAttrs() == 0 {
		// If the record has no Attrs, remove groups at the end of the list; they are empty.
		for len(goas) > 0 && goas[len(goas)-1].group != "" {
			goas = goas[:len(goas)-1]
		}
	}

	indent := 0
	for _, goa := range goas {
		if goa.group != "" {
			buf = fmt.Appendf(buf, "%*s", indent+2, "") // indent
			buf = fmt.Appendf(buf, "%sGroup%s : %s\n", fgColor, colorReset, goa.group)
			indent++
		} else {
			for _, a := range goa.attrs {
				buf = fmt.Appendf(buf, "%*s", 2, "") // indent
				buf = fmt.Appendf(buf, "%sAttrib%s :\n", fgColor, colorReset)
				buf = p.appendAttr(buf, a, fgColor, 1)
				buf = append(buf, '\n')
			}
		}
	}

	// append the attrs if any
	r.Attrs(func(a slog.Attr) bool {
		buf = fmt.Appendf(buf, "%*s", 2, "") // indent
		buf = fmt.Appendf(buf, "%sAttrib%s :\n", fgColor, colorReset)
		buf = p.appendAttr(buf, a, fgColor, 1)
		buf = append(buf, '\n')
		return true
	})

	// add a final new line
	buf = append(buf, '\n')

	// lock, flush, unlock
	p.mu.Lock()
	defer p.mu.Unlock()
	_, err := os.Stderr.Write(buf)

	return err
}

func (p *DevHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return p
	}

	return p.withGroupOrAttrs(groupOrAttrs{attrs: attrs})
}

func (p *DevHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return p
	}

	return p.withGroupOrAttrs(groupOrAttrs{group: name})
}

func (p *DevHandler) withGroupOrAttrs(goa groupOrAttrs) *DevHandler {
	p2 := *p
	p2.goas = make([]groupOrAttrs, len(p.goas)+1)
	copy(p2.goas, p.goas)
	p2.goas[len(p2.goas)-1] = goa

	return &p2
}

func (p *DevHandler) appendLevelMessage(buf []byte, level slog.Level, msg string) ([]byte, []byte, []byte) {
	var colorBg, colorFg []byte

	if level < 0 {
		colorBg = colorBgCyan
		colorFg = colorFgCyan
	} else if level < 4 {
		colorBg = colorBgGreen
		colorFg = colorFgGreen
	} else {
		colorBg = colorBgRed
		colorFg = colorFgRed
	}

	buf = fmt.Appendf(buf, "%s%s %s %s %s%s%s", colorBg, colorFgBlack, level, colorReset, colorFg, msg, colorReset)

	return buf, colorBg, colorFg
}

func (p *DevHandler) appendTime(buf []byte, t time.Time, fgColor []byte) []byte {
	buf = fmt.Appendf(buf, "\n%*s", 2, "") // indent
	buf = fmt.Appendf(buf, "%sTime%s : %s", fgColor, colorReset, t.Format(time.RFC3339))

	return buf
}

func (p *DevHandler) appendSource(buf []byte, pc uintptr, fgColor []byte) []byte {
	f, _ := runtime.CallersFrames([]uintptr{pc}).Next()

	path := f.File
	if len(path) == 0 {
		path = "unknown"
	}

	buf = fmt.Appendf(buf, "\n%*s", 2, "") // indent
	buf = fmt.Appendf(buf, "%sSource%s : %s:%d", fgColor, colorReset, path, f.Line)

	return buf
}

func (p *DevHandler) appendAttr(buf []byte, a slog.Attr, fgColor []byte, indent int) []byte {
	// ignore empty attributes
	if a.Equal(slog.Attr{}) {
		return buf
	}

	a.Value = a.Value.Resolve()

	// in case of error, color the keys red
	valAny := a.Value.Any()
	_, isErr := valAny.(error)
	if isErr == true {
		fgColor = colorFgRed
	}

	// add the attr key
	buf = fmt.Appendf(buf, "%*s", indent*2, "")
	buf = fmt.Appendf(buf, "%s|- %s%s", fgColor, a.Key, colorReset)
	buf = append(buf, " : "...)

	// parse the attr value
	switch a.Value.Kind() {

	case slog.KindString:
		if len(a.Value.String()) > 0 {

			// verify if the string is JSON
			if json.Valid([]byte(a.Value.String())) == true {
				buf = p.appendJSON(buf, []byte(a.Value.String()), indent)
				return buf
			} else {
				buf = append(buf, a.Value.String()...)
			}

		} else {
			buf = fmt.Appendf(buf, "%sempty%s", colorFaint, colorReset)
		}

	case slog.KindTime:
		buf = a.Value.Time().AppendFormat(buf, time.RFC3339)

	// parse groups
	case slog.KindGroup:
		attrs := a.Value.Group()

		// ignore empty groups
		if len(attrs) == 0 {
			return buf
		}

		indent++
		for _, ga := range attrs {
			buf = append(buf, '\n')
			buf = fmt.Appendf(buf, "%*s", 2, "") // indent
			buf = fmt.Appendf(buf, "%sAttrib%s :\n", fgColor, colorReset)
			buf = p.appendAttr(buf, ga, fgColor, indent)
		}

	// it's not a known slog.Kind, figure out the type using reflection
	default:
		valAny := a.Value.Any()

		// Check if valAny is an error
		err, isErr := valAny.(error)
		if isErr == true {
			buf = p.appendError(buf, err, 1)
			return buf
		}

		// append the type, the function resolves the exact types
		buf = p.appendType(buf, reflect.ValueOf(valAny), fgColor, indent)
	}

	return buf
}
