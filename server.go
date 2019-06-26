package main

import (
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	log "github.com/sirupsen/logrus"
	dfpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

type Server struct {
	http.Server
	alive bool
}

func NewServer() *Server {
	s := &Server{
		alive: true,
	}
	return s
}

// Fulfill is the http.HandlerFunc for fulfilling a WebhookRequest,
func (s *Server) Fulfill(w http.ResponseWriter, r *http.Request) {
	var err error
	log.Debugln("ServeHTTP start parse")
	req := &dfpb.WebhookRequest{}
	defer r.Body.Close()
	err = jsonpb.Unmarshal(r.Body, req)
	// err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		log.Errorln("ServeHTTP parse", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Debugln("ServeHTTP start fulfill")
	res, err := s.fulfill(req)
	if err != nil {
		log.Errorln("ServeHTTP fulfill", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Debugln("ServeHTTP start write")
	// err = json.NewEncoder(w).Encode(res)
	err = &jsonpb.Marshaler{}.Marhsal(w, res)
	if err != nil {
		log.Errorln("ServeHTTP write")
	}
}

// Health is the http.HandlerFunc for health checks
func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	if !s.alive {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
