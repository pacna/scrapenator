package main

import (
	"go-image-scraper/internal"
	"os"
)

const (
	// TERMINAL -- enable terminal mode
	TERMINAL = "0"
	// SERVER -- enable server mode
	SERVER = "1"
)

func main() {
	args := os.Args[1:]
	setMode(args)
}

func setMode(args []string) {
	if len(args) > 0 {
		mode := args[0]

		switch mode {
		case TERMINAL:
			internal.InitTerminalMode()
		case SERVER:
			internal.InitServerMode()
		default:
			internal.InitTerminalMode()
		}
	} else {
		internal.InitTerminalMode()
	}
}
