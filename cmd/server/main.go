package main

import (
	"apoor-ssh/pkg/sections"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
)

const (
	defaultHost = "localhost"
	defaultPort = "23234"
)

func main() {
	// Get the config data
	host := os.Getenv("APP_HOST")
	if host == "" {
		host = defaultHost
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = defaultPort
	}

	// Create the server...
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%s", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)
	if err != nil {
		log.Panic(err)
	}

	// Handle interrupts...
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the server...
	log.Printf("Starting SSH server on %s:%s", host, port)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Panic(err)
		}
	}()

	// Wait for interrupt...
	<-done
	log.Println("Shutting down SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		wish.Fatalln(s, "no active terminal, skipping")
		return nil, nil
	}

	m := sections.NewModel().
		WithUserName(s.User()).
		WithWidth(pty.Window.Width).
		WithHeight(pty.Window.Height)
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
