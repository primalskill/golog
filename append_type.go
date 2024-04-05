package golog

import "reflect"

// appendType resolves the value type using reflection and appends it to buf
func (p *DevHandler) appendType(buf []byte, val reflect.Value, fgColor []byte, indent int) []byte {
	kind := val.Type().Kind()

	switch kind {
	case reflect.Bool:
		buf = p.appendPrimitive(buf, val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		buf = p.appendPrimitive(buf, val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		buf = p.appendPrimitive(buf, val)
	case reflect.Float32, reflect.Float64:
		buf = p.appendPrimitive(buf, val)
	case reflect.Complex64, reflect.Complex128:
		buf = p.appendPrimitive(buf, val)
	case reflect.String:
		buf = p.appendString(buf, val)
	case reflect.Array, reflect.Slice:
		buf = p.appendSlice(buf, val, fgColor, indent)
	case reflect.Map:
		buf = p.appendMap(buf, val, fgColor, indent)
	case reflect.Struct:
		buf = p.appendStruct(buf, val, fgColor, indent)
	case reflect.Chan, reflect.Func:
		buf = p.appendUnsupported(buf, kind.String())
		// case reflect.Invalid:
		// 	buf = p.appendInvalid(buf, fieldName, fgColor, indent*2+4)
		// }

	}

	return buf
}
