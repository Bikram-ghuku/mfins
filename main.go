package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type NoticeElement struct {
	SerialNo       int    `json:"slno"`
	MessageId      int    `json:"message_id"`
	MessageSubject string `json:"message_subject"`
	MessageBody    string `json:"message_body"`
	ApprovedOn     string `json:"approved_on"`
	Attachment     int64  `json:"primary_attachemnt_id"` // This is not a spelling mistake by me
	AttachmentURL  string `json:"attachment_url"`
	EventDate      string `json:"event_date"`
	EventTime      string `json:"time_desc"`
	EventVenue     string `json:"event_venue"`
}

var (
	ERPJSession        string
	ERPSSOToken        string
	NoticeEndpoint     string
	FileEndpoint       string
	ERPCatCodeTopicMap map[int]string
	Client             http.Client
	TimeRepeat         int64
	err                error
)

func RunCron() {
	for true {
		log.Println("Getting messages....")
		for key, value := range ERPCatCodeTopicMap {
			log.Printf("Getting notices for %s", value)
			getNewNotices(key)
		}
		time.Sleep(time.Duration(TimeRepeat) * time.Second)
	}
}

func getNewNotices(channel int) {
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
	lastNoticeId := getLastNotice(channel)

	i := 0
	for i < len(resBody) && resBody[i].MessageId != lastNoticeId && resBody[i].SerialNo != lastNoticeId {
		resBody[i].AttachmentURL = fmt.Sprintf(FileEndpoint, resBody[i].Attachment)
		PrintNewMsg(ERPCatCodeTopicMap[channel], resBody[i])
		i++
	}

	if channel > 1000 {
		setLastNotice(channel, resBody[0].SerialNo)
	} else {
		setLastNotice(channel, resBody[0].MessageId)
	}

	if i == 0 {
		log.Printf("No new message on \"%s\"", ERPCatCodeTopicMap[channel])
	}
}

func getLastNotice(channel int) int {
	file, err := os.OpenFile("lastmsg.json", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		log.Panicf("Error opening file: %s", err.Error())
	}

	var fileContent map[int]int
	if err = json.NewDecoder(file).Decode(&fileContent); err != nil {
		log.Panicf("Error decoding file: %s", err.Error())
	}

	return fileContent[channel]
}

func setLastNotice(channel int, lastMsgId int) {
	file, err := os.OpenFile("lastmsg.json", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		log.Panicf("Error opening file: %s", err.Error())
	}

	var fileContent map[int]int
	if err = json.NewDecoder(file).Decode(&fileContent); err != nil {
		log.Panicf("Error decoding file: %s", err.Error())
	}

	fileContent[channel] = lastMsgId

	txt, err := json.Marshal(fileContent)
	if err != nil {
		log.Panicf("Error writing file: %s", err.Error())
	}

	file.Seek(0, 0)
	if _, err = file.Write(txt); err != nil {
		log.Panicf("Error writing file: %s", err.Error())
	}
}

func PrintNewMsg(channel string, content NoticeElement) {
	// this function is called upon receving a new message
	log.Printf("New message on channel %s: \n %v", channel, content.MessageSubject)
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
		Name:  "ssoToken",
		Value: ERPSSOToken,
	})
	return Cookies
}

func main() {
	godotenv.Load()
	ERPJSession = os.Getenv("JSESSIONID")
	ERPSSOToken = os.Getenv("ssoToken")
	TimeRepeat, err = strconv.ParseInt(os.Getenv("REPEAT"), 10, 10)
	if err != nil {
		TimeRepeat = 120
		log.Printf("Error Parsing repeat time, set to 2mins")
	}

	ERPCatCodeTopicMap = map[int]string{
		11:   "Academic",
		12:   "Administrative",
		13:   "Miscellaneous",
		1001: "Academic (UG) section notices",
		1002: "Academic (PG) section notices",
	}

	NoticeEndpoint = "https://erp.iitkgp.ac.in/InfoCellDetails/internal_noticeboard/get_notice_list.htm?cat_code=%d"
	FileEndpoint = "https://erp.iitkgp.ac.in/InfoCellDetails/resources/external/groupemailfile?file_id=%d"

	initClient()

	RunCron()
}
