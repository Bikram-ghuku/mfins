package erplogin

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	rollNo        string
	password      string
	securiyAnswer string
	emailOTP      string
)

func Login(client http.Client) string {
	rollNo = os.Getenv("rollno")
	password = os.Getenv("password")

	fmt.Printf("Security Question: %s", getSecurityQues(&client, rollNo))
	fmt.Println()

	fmt.Printf("Enter answer to security question: ")
	fmt.Scan(&securiyAnswer)

	getOTP(&client)

	fmt.Printf("Enter OTP Sent to your email: ")
	fmt.Scan(&emailOTP)

	loginBody(&client)

	return ""
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

func getOTP(client *http.Client) {
	data := url.Values{}
	data.Set("user_id", rollNo)
	data.Set("password", password)
	data.Set("answer", securiyAnswer)
	data.Set("requestedUrl", HOMEPAGE_URL)
	data.Set("typeee", "SI")

	res, err := client.PostForm(OTP_URL, data)

	if err != nil {
		log.Panic(err.Error())
	}

	dataByte, _ := io.ReadAll(res.Body)

	log.Println(string(dataByte))

	defer res.Body.Close()
}

func loginBody(client *http.Client) {
	data := url.Values{}

	data.Set("user_id", rollNo)
	data.Set("password", password)
	data.Set("answer", securiyAnswer)
	data.Set("email_otp", emailOTP)
	data.Set("requestedUrl", HOMEPAGE_URL)

	resp, err := client.PostForm(LOGIN_URL, data)

	if err != nil {
		log.Panic(err.Error())
	}

	defer resp.Body.Close()
}
