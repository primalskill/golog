package log

import (
	"log/slog"
	"path"
	"runtime"
	"strconv"
	"strings"
)

var (
	_, rootFile, _, _ = runtime.Caller(0)
	rootPath          = path.Dir(rootFile)
)

func prodReplacer(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		a.Key = "t"
	}

	if a.Key == slog.SourceKey {
		s, has := a.Value.Any().(*slog.Source)
		if has == false {
			return a
		}

		s.File = strings.TrimPrefix(s.File, rootPath)
		s.File = strings.TrimPrefix(s.File, "/")

		var sb strings.Builder
		sb.WriteString(s.File)
		sb.WriteString(":")
		sb.WriteString(strconv.Itoa(s.Line))

		a.Value = slog.StringValue(sb.String())
	}

	return a
}
