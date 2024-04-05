package golog

import (
	"fmt"
	"reflect"
)

func (p *DevHandler) appendMap(buf []byte, val reflect.Value, fgColor []byte, indent int) []byte {

	buf = fmt.Appendf(buf, "%s", val.Type().String())

	m := val.MapRange()
	for m.Next() {
		mk := m.Key()
		mv := m.Value()

		buf = append(buf, '\n')
		buf = fmt.Appendf(buf, "%*s", indent+2, "")
		buf = fmt.Appendf(buf, "%s|- ", fgColor)
		buf = p.appendType(buf, mk, fgColor, indent+2)
		buf = fmt.Appendf(buf, "%s : ", colorReset)
		buf = p.appendType(buf, mv, fgColor, indent+2)
	}

	return buf
}
