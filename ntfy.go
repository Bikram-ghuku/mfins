package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/AnthonyHewins/gotfy"
)

func PostData(channel string, content NoticeElement) {
	ctx, cancel := context.WithCancel(context.Background())

	addr := os.Getenv("ntfyAddr")
	server, _ := url.Parse(addr)

	client := http.DefaultClient

	tp, err := gotfy.NewPublisher(server, client)

	if err != nil {
		log.Fatal(err.Error())
	}

	attachURL, _ := url.Parse(content.AttachmentURL)

	if err == nil {
		_, err := tp.SendMessage(ctx, &gotfy.Message{
			Topic:             "test",
			Title:             fmt.Sprintf("#%d | %s | %s | %s", content.MessageId, channel, content.MessageSubject, content.ApprovedOn),
			Message:           content.MessageBody,
			AttachURL:         attachURL,
			AttachURLFilename: "Link 1",
		})

		if err != nil {
			log.Fatal(err.Error())
		}

	} else {
		_, err := tp.SendMessage(ctx, &gotfy.Message{
			Topic:   "test",
			Title:   fmt.Sprintf("#%d | %s | %s", content.MessageId, content.MessageSubject, content.ApprovedOn),
			Message: content.MessageBody,
		})

		if err != nil {
			log.Fatal(err.Error())
		}
	}

	cancel()
}
