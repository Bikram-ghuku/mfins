package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type NoticeElement struct {
	SlNo           string `json:"slno"`
	EventType      string `json:"event_type"`
	MessageSubject string `json:"message_subject"`
	MessageBody    string `json:"message_body"`
	ApprovedOn     string `json:"approved_on"`
	Attachment     int64  `json:"primary_attachemnt_id"`
}

var (
	ERPJSession        string
	ERPSSOToken        string
	NoticeEndpoint     string
	FileEndpoint       string
	erpCatCodeTopicMap map[int]string
)

func MakeRequest(channel int) *http.Request {
	client, err := http.NewRequest("GET", fmt.Sprintf(NoticeEndpoint, channel), nil)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	client.AddCookie(&http.Cookie{
		Name:  "JSESSION",
		Value: ERPJSession,
	})

	client.AddCookie(&http.Cookie{
		Name:  "ssoToken",
		Value: ERPSSOToken,
	})
	return client
}

func main() {
	godotenv.Load()
	ERPJSession = os.Getenv("JSESSIONID")
	ERPSSOToken = os.Getenv("ssoToken")

	erpCatCodeTopicMap = map[int]string{
		11:   "Academic",
		12:   "Administrative",
		13:   "Miscellaneous",
		1001: "Academic (UG) section notices",
		1002: "Academic (PG) section notices",
	}

	NoticeEndpoint = "https://erp.iitkgp.ac.in/InfoCellDetails/internal_noticeboard/get_notice_list.htm?cat_code=%d"
	FileEndpoint = "https://erp.iitkgp.ac.in/InfoCellDetails/resources/external/groupemailfile?file_id=%s"

	client := &http.Client{}
	req := MakeRequest(12)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}

	var resBody []NoticeElement

	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("%v", resBody[0])
}
