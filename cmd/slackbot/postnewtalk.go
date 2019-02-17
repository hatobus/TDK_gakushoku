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
		Pretext: "",
		Text:    "誰かやって~~~~~~~~~~~~~",
		// Uncomment the following part to send a field too
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name: "hoge",
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

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
