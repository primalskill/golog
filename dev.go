package golog

import (
	"fmt"
)

// DevLog represents a development logger, it will use fmt to pretty-print out messages to the console.
type DevLog struct{}

const (
	ColorBlack  string = "\u001b[30m"
	ColorRed           = "\u001b[31m"
	ColorGreen         = "\u001b[32m"
	ColorYellow        = "\u001b[33m"	
	ColorReset         = "\u001b[0m"
)

func NewDevLog() *DevLog {
	return &DevLog{}
}

func (p *DevLog) Info(msg string, meta ...Meta) {
	p.print(msg, "info", meta...)
}

func (p *DevLog) Warn(msg string, meta ...Meta) {
	p.print(msg, "warning", meta...)
}

func (p *DevLog) Error(err error, meta ...Meta) {
	p.print(err.Error(), "error", meta...)
}

func (p *DevLog) print(msg, logType string, meta ...Meta) {
	var hdr, color string

	switch logType {
	case "info":
		hdr = "[INF]"
		color = ColorGreen
	case "warning":
		hdr = "[WARN]"
		color = ColorYellow
	case "error": 
		hdr = "[ERR]"
		color = ColorRed
	}

	fmt.Printf("\n%s%s%s %s\n", color, hdr, ColorReset, msg)

	for _, m := range meta {
		fmt.Printf("  {\n")

		for mk, mv := range m {
			fmt.Printf("    %s%s:%s %+v\n", color, mk, ColorReset, mv)			
		}		
		fmt.Printf("  }\n")
	}
	
	fmt.Printf("\n")
}







