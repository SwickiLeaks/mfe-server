package main

import (
	"mfeserver/internal/server"
)

// Main Function
func main() {
	// logger := log.New(os.Stdout, "INFO: ", log.Lshortfile|log.Ldate|log.Ltime)

	// config, err := config.LoadAppShellConfig()
	// if err != nil {
	// 	log.Fatal("Error parsing appshell configs")
	// }

	// TODO: Process appshell config

	srv := server.NewServer()
	srv.Run(":4321")
}
