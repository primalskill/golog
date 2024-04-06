package golog

import (
	"fmt"
	"strings"

	"github.com/primalskill/errors"
)

// appendError takes err and converts it to error.Error then appends it to buf parsing with appendPskError.
func (p *DevHandler) appendError(buf []byte, err error, indent int) []byte {
	if err == nil {
		return buf
	}

	// convert and flatten err (if it's wrapped) to []errors.Error
	errs := errors.Flatten(err)

	for i, err := range errs {
		if i+1 < len(errs) {
			buf = p.appendPskError(buf, err, errs[i+1], indent)
		} else {
			buf = p.appendPskError(buf, err, errors.Error{}, indent)
		}
	}

	return buf
}

// appendPskError adds err to buf where err is of errors.Error. nextErr is passed to the function to trim the wrapped
// error messages.
func (p *DevHandler) appendPskError(buf []byte, err, nextErr errors.Error, indent int) []byte {
	if len(err.Msg) == 0 {
		return buf
	}

	msgError := err.Error()
	msgNextError := nextErr.Error()

	// trim msgNextError from msgError, because Go will add every wrapped error message to the current error message
	msg, _ := strings.CutSuffix(msgError, msgNextError)
	msg = strings.TrimSpace(msg)
	msg = strings.TrimRight(msg, ":")

	// add the message
	buf = append(buf, '\n')
	buf = fmt.Appendf(buf, "%*s", indent*2+3, "")
	buf = fmt.Appendf(buf, "%s|- error%s : %s", colorFgRed, colorReset, msg)

	// add the source if any
	if len(err.Source) > 0 {
		buf = append(buf, '\n')
		buf = fmt.Appendf(buf, "%*s", indent*2+3, "")
		buf = fmt.Appendf(buf, "%s|- source%s : %s", colorFgRed, colorReset, err.Source)
	}

	// output the err Meta if any
	if len(err.Meta) > 0 {
		buf = append(buf, '\n')
		buf = fmt.Appendf(buf, "%*s", indent*2+3, "")
		buf = fmt.Appendf(buf, "%s|- meta%s", colorFgRed, colorReset)

		for mk, mv := range err.Meta {
			buf = append(buf, '\n')
			buf = fmt.Appendf(buf, "%*s", indent*2+5, "")
			buf = fmt.Appendf(buf, "%s|- %s%s : %+v", colorFgRed, mk, colorReset, mv)
		}
	}

	return buf
}
