package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server Structure
type Server struct {
	server *http.Server
	router *mux.Router
}

// Create New Server
func NewServer() *Server {
	s := Server{
		server: &http.Server{
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  5 * time.Second,
		},
		router: mux.NewRouter().StrictSlash(true),
	}

	s.routes()

	s.server.Handler = handlers.CORS()(s.router)
	return &s
}

// Run Server
func (s *Server) Run(port string) {
	ctx, cancel := context.WithCancel(context.Background())
	s.server.BaseContext = func(_ net.Listener) context.Context { return ctx }

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	s.server.Addr = port

	go func() {
		log.Println("Server starting on port: ", s.server.Addr)
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal("HTTP Server Error")
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
	x := <-signalChan
	log.Printf("os.Interrupt %s - shutting down...\n", x)

	// Terminate after second signal before callback is done
	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...")
	}()

	// Perform graceful shutdown here
	gracefulCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := s.server.Shutdown(gracefulCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
	}

	log.Printf("gracefully stopped\n")
	cancel()

	defer os.Exit(0)
}
