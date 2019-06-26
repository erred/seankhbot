package main

import (
	log "github.com/sirupsen/logrus"
	dfpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// fulfill actually fulfills a WebhookRequest
func (s *Server) fulfill(r *dfpb.WebhookRequest) (*dfpb.WebhookResponse, error) {
	log.Debugln("fulfill started with", r)
	res := &dfpb.WebhookResponse{
		FulfillmentText: "I am fulfilled!!!",
	}
	return res, nil
}
