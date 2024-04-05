package golog

import (
	"fmt"
	"reflect"
)

func (p *DevHandler) appendSlice(buf []byte, val reflect.Value, fgColor []byte, indent int) []byte {

	buf = fmt.Appendf(buf, "%s", val.Type().String())

	for i := 0; i < val.Len(); i++ {
		buf = append(buf, '\n')
		buf = fmt.Appendf(buf, "%*s", indent+2, "")
		buf = fmt.Appendf(buf, "%s|- %d%s : ", fgColor, i, colorReset)

		buf = p.appendType(buf, val.Index(i), fgColor, indent+2)
	}

	return buf
}
