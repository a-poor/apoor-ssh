package main

import (
	"apoor-ssh/pkg/sections"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
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

	// Create the logger...
	logger := log.New(
		log.WithOutput(os.Stdout),
		log.WithLevel(log.DebugLevel),
		log.WithTimestamp(),
		log.WithCaller(),
		log.WithPrefix("apoor-dot-ssh"),
	)
	logger.Debug("Starting...")

	// Create the tea handler...
	teaHandler := makeTeaHandler(logger)

	// Create the server...
	logger.Debug("Creating SSH server...")
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%s", host, port)),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)
	if err != nil {
		logger.Fatal(err)
	}

	// Handle interrupts...
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the server...
	logger.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			logger.Fatal("ListenAndServe completed.", "err", err)
		}
	}()

	// Wait for interrupt...
	<-done
	logger.Info("Gracefully shutting down SSH server (with 30s timeout)")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Error from shutdown.", "err", err)
	}
}

func makeTeaHandler(logger log.Logger) func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	return func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
		logger.Debug("New ssh session checking pty...")
		pty, _, active := s.Pty()
		logger.Debug(
			"New ssh session",
			"term", pty.Term,
			"active", active,
		)
		if !active {
			logger.Error("No active terminal, skipping")
			wish.Fatalln(s, "no active terminal, skipping")
			return nil, nil
		}

		m := sections.NewModel().
			WithUserName(s.User()).
			WithWidth(pty.Window.Width).
			WithHeight(pty.Window.Height).
			WithLogger(logger)

		return m, []tea.ProgramOption{tea.WithAltScreen()}
	}
}
