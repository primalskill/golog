package golog

import (
	"fmt"
	"reflect"
)

// appendPrimitive adds primitive types that can be converted to interface{} to buf.
func (p *DevHandler) appendPrimitive(buf []byte, val reflect.Value) []byte {
	if val.CanInterface() == false {
		buf = fmt.Appendf(buf, "%sinvalid type%s", colorFgRed, colorReset)
		return buf
	}

	buf = fmt.Appendf(buf, "%v", val)
	return buf
}

// appendString adds val as string to buf. val can be empty, in this case write out empty.
func (p *DevHandler) appendString(buf []byte, val reflect.Value) []byte {
	str := val.String()

	if len(str) == 0 {
		buf = fmt.Appendf(buf, "%sempty%s", colorFaint, colorReset)
		return buf
	}

	buf = fmt.Appendf(buf, "%s", str)
	return buf
}

// appendUnsupported adds unsupported as string to buf. func and chan types are unsupported.
func (p *DevHandler) appendUnsupported(buf []byte, kind string) []byte {
	buf = fmt.Appendf(buf, "%s%s (unsupported)%s", colorFaint, kind, colorReset)
	return buf
}

// appendInvalid adds "invalid type" to buf.
func (p *DevHandler) appendInvalid(buf []byte) []byte {
	buf = fmt.Appendf(buf, "%sinvalid type%s", colorFaint, colorReset)
	return buf
}

// appendNil adds "nil" to buf.
func (p *DevHandler) appendNil(buf []byte) []byte {
	buf = fmt.Appendf(buf, "%snil%s", colorFaint, colorReset)
	return buf
}
