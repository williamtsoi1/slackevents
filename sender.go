package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	ev "github.com/mchmarny/gcputil/env"
	"github.com/nlopes/slack"
)

var (
	token   = ev.MustGetEnvVar("SLACK_API_TOKEN", "")
	channel = ev.MustGetEnvVar("SLACK_CHANNEL", "")
)

// Send sends the message
func send(tw *twitter.Tweet) error {

	if tw == nil {
		return fmt.Errorf("Null tweet on send")
	}

	api := slack.New(token)

	a1 := slack.Attachment{
		Title:     fmt.Sprintf("%s (%s)", tw.User.ScreenName, tw.User.Name),
		TitleLink: fmt.Sprintf("https://twitter.com/%s", tw.User.IDStr),
		ImageURL:  tw.User.ProfileImageURL,
	}

	a1.Fields = []slack.AttachmentField{
		slack.AttachmentField{
			Title: "Tweet:",
			Value: tw.Text,
		},
		slack.AttachmentField{
			Title: "Link:",
			Value: fmt.Sprintf("https://twitter.com/%s/status//%s", tw.User.IDStr, tw.IDStr),
		},
	}

	_, _, err := api.PostMessage(channel,
		slack.MsgOptionText("Twitter", false),
		slack.MsgOptionAttachments(a1))

	log.Printf("Sent tweet: %s", tw.IDStr)

	return err

}
