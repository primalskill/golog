package golog

import (
	"fmt"
	"reflect"
)

// appendStruct takes val and appends it to buf as a struct.
func (p *DevHandler) appendStruct(buf []byte, val reflect.Value, fgColor []byte, indent int) []byte {

	vType := val.Type()

	buf = fmt.Appendf(buf, "%s", vType.String())

	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name

		buf = append(buf, '\n')
		buf = fmt.Appendf(buf, "%*s", indent+2, "")
		buf = fmt.Appendf(buf, "%s|- %s%s : ", fgColor, fieldName, colorReset)
		buf = p.appendType(buf, val.Field(i), fgColor, indent+2)
	}

	return buf
}
