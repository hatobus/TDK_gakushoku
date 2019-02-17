package slackbot

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

func PostNewTalk(userID, category string) error {
	api := slack.New(os.Getenv("SlackTOKEN"))

	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	channelID, timestamp, err := api.PostMessage(os.Getenv("slackChannelID"), slack.MsgOptionText(category, false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
