package golog

import "reflect"

func (p *DevHandler) appendInterface(buf []byte, val reflect.Value) []byte {

	// check if val implements the error interface handle it as an error
	errInterface := reflect.TypeOf((*error)(nil)).Elem()
	if val.Type().Implements(errInterface) {
		err, isErr := val.Interface().(error)
		if isErr == true {
			buf = p.appendError(buf, err, 1)
			return buf
		}
	}

	buf = p.appendPrimitive(buf, val)
	return buf
}
