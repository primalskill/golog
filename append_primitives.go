package golog

import (
	"fmt"
	"reflect"
)

func (p *DevHandler) appendPrimitive(buf []byte, val reflect.Value) []byte {
	if val.CanInterface() == false {
		buf = fmt.Appendf(buf, "%sinvalid type%s", colorFgRed, colorReset)
		return buf
	}

	buf = fmt.Appendf(buf, "%v", val)
	return buf
}

func (p *DevHandler) appendString(buf []byte, val reflect.Value) []byte {
	str := val.String()

	if len(str) == 0 {
		buf = fmt.Appendf(buf, "%sempty%s", colorFaint, colorReset)
		return buf
	}

	buf = fmt.Appendf(buf, "%s", str)
	return buf
}

func (p *DevHandler) appendUnsupported(buf []byte, kind string) []byte {
	buf = fmt.Appendf(buf, "%s%s (unsupported)%s", colorFaint, kind, colorReset)
	return buf
}
