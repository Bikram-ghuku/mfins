package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ntfyMsg struct {
	Title    string
	Body     string
	Tags     []string
	Priority int
	Link     string
	Filename string
}

func PostData(channel string, content NoticeElement) {

	postNtfy(Addr, "test", ntfyMsg{
		Title:    fmt.Sprintf("#%d | %s | %s | %s", content.SerialNo, content.MessageSubject, channel, content.ApprovedOn),
		Body:     content.MessageBody,
		Priority: 5,
		Link:     content.AttachmentURL,
		Filename: "Link 1",
		Tags:     []string{"Warning", "cd"},
	})
}

func postNtfy(addr string, channel string, msg ntfyMsg) {

	body := fmt.Sprintf(`{
		"topic": "%s",
		"message": "%s",
		"title": "%s",
		"priority": %d,
		"attach": "%s",
		"filename": "%s"
	}`, channel, msg.Body, msg.Title, msg.Priority, msg.Link, msg.Filename)

	req, _ := http.NewRequest("POST", addr, strings.NewReader(body))

	if NtfyUser != "" && NtfyPass != "" {
		authHeader := fmt.Sprintf("Basic %s:%s", NtfyUser, NtfyPass)
		req.Header.Set("Authorization", authHeader)
	}

	req.Header.Set("Markdown", "yes")

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error publishing to ntfy: %s", err)
	}
}
