package main

import (
	log "github.com/sirupsen/logrus"
	dfpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// fulfill actually fulfills a WebhookRequest
func (s *Server) fulfill(r *dfpb.WebhookRequest) (*dfpb.WebhookResponse, error) {
	log.Debugln("fulfill session", r.Session)
	if r.QueryResult != nil {
		log.Debugln("fulfill QR", r.QueryResult.Action, r.QueryResult.QueryText)
		log.Debugf("fulfill QR intent %v\n", r.QueryResult.Intent)

	}
	if r.OriginalDetectIntentRequest != nil {
		log.Debugln("fulfill ODIR", r.OriginalDetectIntentRequest.Source)
		log.Debugf("fulfill ODIR payload %v\n", r.OriginalDetectIntentRequest.Payload)
	}
	res := &dfpb.WebhookResponse{
		FulfillmentText: "I am fulfilled!!!",
	}
	return res, nil
}
