package main

import (
	"embed"
	"log"
	"mfeserver/config"
	"mfeserver/internal/server"
	"os"
)

//go:embed appshell-config.json
var configFS embed.FS

// Server entry
func main() {
	// Logger that writes to STDOUT
	logger := log.New(os.Stdout, "INFO: ", log.Lshortfile|log.Ldate|log.Ltime)

	// Load and parse appshell config
	logger.Println("Loading Appshell Configuration...")
	config, err := config.LoadAppShellConfig()
	if err != nil {
		log.Fatal("Error parsing appshell configs")
	}

	// TODO: Process appshell config
	for shell, shellContents := range config.Appshells {
		logger.Println(shell)
		for k, v := range shellContents {
			logger.Println("    ", k, v)
		}
	}

	// TODO: Write Appshell Config to the template

	// TODO: Setup server routes and handlers
	srv := server.NewServer()
	srv.Run(":4321")
}
