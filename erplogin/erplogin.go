package erplogin

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type loginDetails struct {
	user_id      string
	password     string
	answer       string
	requestedUrl string
	email_otp    string
}

type erpCreds struct {
	RollNumber               string            `json:"roll_number"`
	Password                 string            `json:"password"`
	SecurityQuestionsAnswers map[string]string `json:"answers"`
}

func request_otp(client http.Client, loginParams loginDetails) {
	data := url.Values{}
	data.Set("typeee", "SI")
	data.Set("user_id", loginParams.user_id)
	data.Set("password", loginParams.password)
	data.Set("answer", loginParams.answer)

	res, err := client.PostForm(OTP_URL, data)
	if err != nil {
		log.Printf(err.Error())
	}

	defer res.Body.Close()
}

func getCreds(client *http.Client) loginDetails {
	loginParams := loginDetails{
		requestedUrl: HOMEPAGE_URL,
	}

	fmt.Print("Enter Roll No.: ")
	fmt.Scan(&loginParams.user_id)

	fmt.Print("Enter ERP Password: ")
	byte_password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Printf(err.Error())
	}

	fmt.Printf("Your secret question: %s\n", getSecretQuestion(client, loginParams.user_id))
	fmt.Print("Enter answer to your secret question: ")
	byte_answer, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Printf(err.Error())
	}

	fmt.Println()

	loginParams.answer = string(byte_answer)
	loginParams.password = string(byte_password)

	request_otp(*client, loginParams)

	fmt.Println("Enter your OTP sent to email: ")
	byte_otp, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Printf(err.Error())
	}

	loginParams.email_otp = string(byte_otp)

	return loginParams
}

func getSecretQuestion(client *http.Client, roll_number string) string {
	data := map[string][]string{
		"user_id": {roll_number},
	}

	res, err := client.PostForm(SECRET_QUESTION_URL, data)
	if err != nil {
		log.Printf(err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf(err.Error())
	}

	return string(body)
}

func ERPSession() string {
	client := http.Client{}
	loginParams := getCreds(&client)
	data := url.Values{}
	data.Set("user_id", loginParams.user_id)
	data.Set("password", loginParams.password)
	data.Set("answer", loginParams.answer)
	data.Set("requestedUrl", loginParams.requestedUrl)
	data.Set("email_otp", loginParams.email_otp)

	res, err := client.PostForm(LOGIN_URL, data)
	if err != nil {
		log.Printf(err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf(err.Error())
	}

	bodyStr := string(body)
	idx := strings.Index(bodyStr, "ssoToken")
	ssoToken := bodyStr[strings.LastIndex(bodyStr[:idx], "\"")+1 : strings.Index(bodyStr, "ssoToken")+strings.Index(bodyStr[idx:], "\"")]

	return ssoToken

}
