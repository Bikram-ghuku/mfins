package erplogin

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	rollNo         string
	password       string
	securityAnswer string
	emailOTP       string
)

func Login(client http.Client) {

	if checkLogin(&client) {
		log.Println("Already logged in, continuing without login")
		return
	}

	rollNo = os.Getenv("ROLL_NO")
	password = os.Getenv("PASSWORD")

	question := getSecurityQues(&client, rollNo)
	securityAnswer = GetSecurityAnswer(question)

	emailOTP = GetOtp(&client)

	loginBody(&client)
}

func getSecurityQues(client *http.Client, rollNo string) string {
	data := url.Values{}

	data.Set("user_id", rollNo)

	res, err := client.PostForm(SECRET_QUESTION_URL, data)

	if err != nil {
		log.Panic(err.Error())
	}

	byteResp, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	return string(byteResp)
}

func SendOTP(client *http.Client) {
	log.Println("Requesting OTP....")
	data := url.Values{}
	data.Set("user_id", rollNo)
	data.Set("password", password)
	data.Set("answer", securityAnswer)
	data.Set("requestedUrl", HOMEPAGE_URL)
	data.Set("typeee", "SI")

	res, err := client.PostForm(OTP_URL, data)

	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("OTP requested successfully!")

	defer res.Body.Close()
}

func loginBody(client *http.Client) {
	data := url.Values{}

	data.Set("user_id", rollNo)
	data.Set("password", password)
	data.Set("answer", securityAnswer)
	data.Set("email_otp", emailOTP)
	data.Set("requestedUrl", HOMEPAGE_URL)

	resp, err := client.PostForm(LOGIN_URL, data)

	if err != nil {
		log.Panic(err.Error())
	}

	defer resp.Body.Close()
}

func checkLogin(client *http.Client) bool {

	res, err := client.Get(WELCOMEPAGE_URL)

	if err != nil {
		log.Panic(err.Error())
	}

	return res.ContentLength == 1034
}
