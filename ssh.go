package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"os"
	"os/signal"
	"syscall"
	"time"

	gossh "golang.org/x/crypto/ssh"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
)

const (
	host = "0.0.0.0"
	port = 4202
)

type Session struct {
	Id string
}

type User struct {
	Name      string
	PublicKey string
}

func sshHandler(h ssh.Handler) ssh.Handler {
	return func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello, %s!\n", s.User()))
		authorizedKey := gossh.MarshalAuthorizedKey(s.PublicKey())
		io.WriteString(s, fmt.Sprintf("You used this public key to authenticate:\n%s", authorizedKey))
		d := User{Name: s.User(), PublicKey: strings.TrimSpace(string(authorizedKey))}
		file, _ := json.MarshalIndent(d, "", " ")
		os.WriteFile(fmt.Sprintf("/tmp/%s", s.Command()[0]), file, 0644)
		h(s)
	}
}

func SSHServe() {
	pem := []byte(os.Getenv("HOST_PRIVATE_KEY"))

	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPEM(pem),
		wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
			return true
		}),
		wish.WithMiddleware(
			sshHandler,
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("could not start server", "error", err)
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("could not start server", "error", err)
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("could not stop server", "error", err)
	}
}
