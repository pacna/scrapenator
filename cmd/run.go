package cmd

import (
	"flag"

	"github.com/pacna/scrapenator/internal/server"
	"github.com/pacna/scrapenator/internal/terminal"
)

func Execute() {
	mode := flag.String("mode", "server", "allow user to switch mode (server or terminal)")

	flag.Parse()

	switch *mode {
	case "terminal":
		terminal.New()
	case "server":
		fallthrough
	default:
		server.New()
	}
}