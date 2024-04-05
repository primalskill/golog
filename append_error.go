package golog

import "github.com/primalskill/errors"

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
