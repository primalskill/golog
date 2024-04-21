package golog

import (
	"fmt"
	"reflect"
)

func (p *DevHandler) appendSlice(buf []byte, val reflect.Value, fgColor []byte, indent int) []byte {

	// If val implements fmt.Stringer interface, than call it.
	// This is useful for types like UUID

	stringer := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	if val.Type().Implements(stringer) == true {
		valMethod := val.MethodByName("String")

		// Double check if the method value returned is valid
		if valMethod.Kind() != reflect.Invalid {
			retVal := valMethod.Call([]reflect.Value{})

			buf = append(buf, retVal[0].String()...)
			return buf
		}
	}

	// []byte is an alias for []uint8 which is a slice, test if it's truly []byte or not
	if val.CanConvert(reflect.TypeOf([]byte(nil))) {
		b := val.Bytes()
		buf = append(buf, string(b)...)

		return buf
	}

	// Handle it like a regular slice
	buf = fmt.Appendf(buf, "%s", val.Type().String())

	for i := 0; i < val.Len(); i++ {
		buf = append(buf, '\n')
		buf = fmt.Appendf(buf, "%*s", indent+2, "")
		buf = fmt.Appendf(buf, "%s|- %d%s : ", fgColor, i, colorReset)

		buf = p.appendType(buf, val.Index(i), fgColor, indent+2)
	}

	return buf
}
