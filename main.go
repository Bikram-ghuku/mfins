package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type NoticeElement struct {
	SerialNo       int    `json:"slno"`
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

func RunCron() {
	for true {
		log.Println("Getting messages....")
		for key, value := range erpCatCodeTopicMap {
			log.Printf("Getting notices for %s", value)
			getNotices(key)
		}
		time.Sleep(5 * time.Second)
	}
}

func getNotices(channel int) {
	req, err := http.NewRequest("GET", fmt.Sprintf(NoticeEndpoint, channel), nil)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	resp, err := Client.Do(req)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	var resBody []NoticeElement

	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	if channel < 1000 {
		log.Printf("Last message id: %d", resBody[0].MessageId)
	} else {
		log.Printf("Last message id: %d", resBody[0].SerialNo)
	}
}

func initClient() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	parseURL, _ := url.Parse(NoticeEndpoint)

	jar.SetCookies(parseURL, MakeCookies())

	Client = http.Client{
		Jar: jar,
	}
}

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

	initClient()

	RunCron()
}
