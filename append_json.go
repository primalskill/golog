package golog

import (
	"fmt"

	"github.com/tidwall/pretty"
)

var jsonColorStyle = &pretty.Style{
	Key:      [2]string{"\x1B[95m", "\x1B[0m"},
	String:   [2]string{"\x1B[32m", "\x1B[0m"},
	Number:   [2]string{"\x1B[33m", "\x1B[0m"},
	True:     [2]string{"\x1B[36m", "\x1B[0m"},
	False:    [2]string{"\x1B[36m", "\x1B[0m"},
	Null:     [2]string{"\x1B[2m", "\x1B[0m"},
	Escape:   [2]string{"\x1B[35m", "\x1B[0m"},
	Brackets: [2]string{"\x1B[0m", "\x1B[0m"},
	Append:   pretty.TerminalStyle.Append,
}

func (p *DevHandler) appendJSON(buf, jsonData []byte, indent int) []byte {
	var jsonPrettyOpts = &pretty.Options{
		Width:    80,
		Prefix:   fmt.Sprintf("%*s", indent+2, " "),
		Indent:   "  ",
		SortKeys: false,
	}

	jsonData = pretty.PrettyOptions(jsonData, jsonPrettyOpts)
	jsonData = pretty.Color(jsonData, jsonColorStyle)

	buf = append(buf, "JSON"...)
	buf = append(buf, '\n')
	buf = append(buf, jsonData...)

	return buf
}
