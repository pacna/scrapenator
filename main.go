package main

import (
	"go-image-scraper/utils"
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
			utils.InitTerminalMode()
		case SERVER:
			utils.InitServerMode()
		default:
			utils.InitTerminalMode()
		}
	} else {
		utils.InitTerminalMode()
	}
}
