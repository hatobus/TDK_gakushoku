package slackbot

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func PostNewTalk(userID, category string) error {
	api := slack.New(os.Getenv("NAMIKI_SlackTOKEN"))
	log.Println(os.Getenv("NAMIKI_SlackTOKEN"))

	attachment := slack.Attachment{
		Pretext:    "",
		Text:       "誰かやって~~~~~~~~~~~~~",
		Fallback:   "who are doing Namiki's work?",
		CallbackID: "user/request",
		// Uncomment the following part to send a field too
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name: "request",
				Text: "俺がやる",
				Type: "button",
			},
		},
	}

	channelID, timestamp, err := api.PostMessage(
		os.Getenv("slackChannelID"),
		slack.MsgOptionText(userID+" さんが "+category+" の仕事を誰かにやってもらいたいそうです", false),
		slack.MsgOptionAttachments(attachment))

	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
