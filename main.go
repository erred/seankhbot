package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var (
	Port = os.Getenv("PORT")
)

func init() {
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "ERROR":
		fallthrough
	default:
		log.SetLevel(log.ErrorLevel)
	}
	log.Debugln("Log level set to", log.GetLevel())

	switch os.Getenv("LOG_FORMAT") {
	case "JSON":
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
