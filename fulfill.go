package main

import (
	dfpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

// fulfill actually fulfills a WebhookRequest
func (s *Server) fulfill(r *dfpb.WebhookRequest) (*dfpb.WebhookResponse, error) {
	res := &dfpb.WebhookResponse{
		FulfillmentText: "I am fulfilled!!!",
	}
	return res, nil
}
