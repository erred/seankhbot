package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var (
	Port = os.Getenv("PORT")
)

func init() {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		fallthrough
	default:
		log.SetLevel(log.ErrorLevel)
	}
	log.Debugln("Log level set to", log.GetLevel())

	switch strings.ToLower(os.Getenv("LOG_FORMAT")) {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
		log.Debugln("Log format set to json")
	default:
		log.SetFormatter(&log.TextFormatter{})
		log.Debugln("Log format set to text")
	}

	if Port == "" {
		Port = ":8080"
	}
}

func main() {
	ctx := context.Background()

	log.Debugln("main create server")
	svr := NewServer()

	s := &http.Server{
		Addr: Port,
	}
	defer s.Shutdown(ctx)

	http.HandleFunc("/health", svr.Health)
	http.HandleFunc("/", svr.Fulfill)

	log.Infoln("Starting on port", s.Addr)
	go s.ListenAndServe()

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGKILL)
	sig := <-sigs
	log.Debugln("main ending with", sig)
}
