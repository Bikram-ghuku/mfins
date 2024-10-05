package erplogin

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

const (
	query       = "from:erpkgp@adm.iitkgp.ac.in is:unread subject: otp"
	RedirectURL = "http://localhost:7007"
)

func GetOtp(client *http.Client) string {
	if is_file("client_secret.json") || is_file(".token") {
		return GetOtpEmail(client)
	} else {
		return GetOtpInput(client)
	}
}

func getMsgId(service *gmail.Service) string {
	results, err := service.Users.Messages.List("me").Q(query).MaxResults(1).Do()

	if err != nil {
		log.Fatal(err.Error())
	}

	if len(results.Messages) != 0 {
		return results.Messages[0].Id
	}

	return ""

}

func GetOtpEmail(client *http.Client) string {
	ctx, cancel := context.WithCancel(context.Background())

	conf := oauth2.Config{
		Scopes:      []string{gmail.GmailReadonlyScope},
		Endpoint:    google.Endpoint,
		RedirectURL: RedirectURL,
	}

	secretByte, err := os.ReadFile("client_secret.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	var secret map[string]map[string]json.RawMessage
	err = json.Unmarshal(secretByte, &secret)

	_ = json.Unmarshal(secret["installed"]["client_id"], &conf.ClientID)
	_ = json.Unmarshal(secret["installed"]["client_secret"], &conf.ClientSecret)

	var token *oauth2.Token

	if is_file(".token") {

		token_byte, err := os.ReadFile(".token")
		if err != nil {
			log.Fatal(err.Error())
		}

		err = json.Unmarshal(token_byte, &token)
		if err != nil {
			log.Fatal(err.Error())
		}

	} else {
		token, err = generateToken(&ctx, cancel, &conf)
		if err != nil {
			log.Fatal(err.Error())
		}

		token_json, err := json.Marshal(*token)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = os.WriteFile(".token", token_json, 0666)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	service, err := gmail.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, token)))
	if err != nil {
		log.Fatal(err.Error())
	}

	latestId := getMsgId(service)
	SendOTP(client)
	var mailId string

	for {
		log.Println("Waiting for OTP...")
		if mailId = getMsgId(service); mailId != latestId {
			log.Println("OTP fetched")
			break
		}
		time.Sleep(1 * time.Second)
	}

	message, err := service.Users.Messages.Get("me", mailId).Do()
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := base64.URLEncoding.DecodeString(message.Payload.Body.Data)
	if err != nil {
		log.Fatal(err.Error())
	}

	reg := regexp.MustCompile("[0-9]+")
	otp := reg.FindAllString(string(body), -1)[0]

	cancel()
	return otp

}

func is_file(filename string) bool {
	file, err := os.Open(filename)
	defer file.Close()
	return !errors.Is(err, os.ErrNotExist)
}

func generateToken(ctx *context.Context, cancel context.CancelFunc, conf *oauth2.Config) (*oauth2.Token, error) {
	authURL := conf.AuthCodeURL("666573902ekwjfn")
	fmt.Println("Visit this URL for authentication: ", authURL)
	browser.OpenURL(authURL)

	var token *oauth2.Token
	var err error

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") == "666573902ekwjfn" {
			token, err = conf.Exchange(*ctx, r.URL.Query().Get("code"))
		}
		fmt.Fprintf(w, "Authentication complete. Check your terminal.")
		cancel()
	})

	server := &http.Server{Addr: ":7007"}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	<-(*ctx).Done()
	return token, err
}

func GetOtpInput(client *http.Client) string {
	SendOTP(client)
	var emailOTP string
	fmt.Printf("Enter OTP Sent to your email: ")
	fmt.Scan(&emailOTP)

	return emailOTP
}
