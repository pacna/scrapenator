package cmd

import (
	"flag"

	"github.com/pacna/scrapenator/internal/terminal"
)

func Execute() {
	mode := flag.String("mode", "server", "allow user to terminal")

	flag.Parse()

	switch *mode {
	case "terminal":
		terminal.New()
	default:
		terminal.New()
	}
}