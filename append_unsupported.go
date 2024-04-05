package golog

import "fmt"

// func (p *DevHandler) appendUnsupported(buf []byte, name string, kind string, fgColor []byte, indent int) []byte {
// 	buf = fmt.Appendf(buf, "%*s", indent, "")
// 	buf = fmt.Appendf(buf, "%s%s%s : %s %s(unsupported)%s", fgColor, name, colorReset, kind, colorFaint, colorReset)

// 	return buf
// }

func (p *DevHandler) appendInvalid(buf []byte, name string, fgColor []byte, indent int) []byte {
	buf = fmt.Appendf(buf, "%*s", indent, "")
	buf = fmt.Appendf(buf, "%s%s%s : %s(invalid)%s", fgColor, name, colorReset, colorFgRed, colorReset)

	return buf
}
