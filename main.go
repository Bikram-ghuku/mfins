package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type NoticeElement struct {
	MessageId      int    `json:"message_id"`
	MessageSubject string `json:"message_subject"`
	MessageBody    string `json:"message_body"`
	ApprovedOn     string `json:"approved_on"`
	Attachment     int64  `json:"primary_attachemnt_id"`
	EventDate      string `json:"event_date"`
	EventTime      string `json:"time_desc"`
	EventVenue     string `json:"event_venue"`
}

var (
	ERPJSession        string
	ERPSSOToken        string
	NoticeEndpoint     string
	FileEndpoint       string
	erpCatCodeTopicMap map[int]string
	Client             http.Client
)

func MakeCookies() []*http.Cookie {

	var Cookies []*http.Cookie
	Cookies = append(Cookies, &http.Cookie{
		Name:  "JSESSION",
		Value: ERPJSession,
	})

	Cookies = append(Cookies, &http.Cookie{
		Name:  "ssoToken",
		Value: ERPSSOToken,
	})
	return Cookies
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

	req, err := http.NewRequest("GET", fmt.Sprintf(NoticeEndpoint, 11), nil)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	parseURL, _ := url.Parse(NoticeEndpoint)

	jar.SetCookies(parseURL, MakeCookies())

	Client = http.Client{
		Jar: jar,
	}

	resp, err := Client.Do(req)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	var resBody []NoticeElement

	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	fmt.Printf("%v", resBody[0])
}
