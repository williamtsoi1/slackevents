package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	ce "github.com/cloudevents/sdk-go"
	"github.com/dghubble/go-twitter/twitter"
)

type eventReceiver struct{}

func (r *eventReceiver) Receive(ctx context.Context, event ce.Event, resp *ce.EventResponse) error {

	// get content

	if event.DataContentType() != "application/json" {
		return fmt.Errorf("Invalid Data Content Type: %s. Only application/json supported",
			event.DataContentType())
	}

	content, err := event.DataBytes()
	if err != nil {
		log.Printf("Failed to DataAs bytes: %s", err.Error())
		return err
	}

	var tw twitter.Tweet
	if err := json.Unmarshal(content, &tw); err != nil {
		log.Printf("Error while decoding tweet: %s", err.Error())
		return err
	}

	log.Printf("Tweet to post: %s", tw.IDStr)
	return send(&tw)

}
