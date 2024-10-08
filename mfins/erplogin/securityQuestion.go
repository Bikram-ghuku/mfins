package erplogin

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var questionAnswer map[string]string

const FILENAME = "security_question.json"

func GetSecurityAnswer(question string) string {

	if is_file(FILENAME) {

		log.Printf("Found %s, checking if question is present....", FILENAME)

		quesByte, err := os.ReadFile(FILENAME)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = json.Unmarshal(quesByte, &questionAnswer)
		if err != nil {
			log.Fatal(err.Error())
		}

		if val, ok := questionAnswer[question]; ok {
			log.Println("Found answer to security question!")
			return val
		} else {
			inputAnswer(question)
		}

		questionAnswerJson, err := json.Marshal(questionAnswer)
		if err != nil {
			log.Fatalf(err.Error())
		}

		os.WriteFile(FILENAME, questionAnswerJson, 0666)

		return questionAnswer[question]

	} else {
		log.Printf("%s not found, creating.....", FILENAME)
		inputAnswer(question)

		questionAnswerJson, err := json.Marshal(questionAnswer)
		if err != nil {
			log.Fatalf(err.Error())
		}

		os.WriteFile(FILENAME, questionAnswerJson, 0666)

		return questionAnswer[question]
	}
}

func inputAnswer(question string) {
	var answer string
	fmt.Printf("Security question: %s", question)
	fmt.Println()
	fmt.Print("Enter answer to security question: ")
	fmt.Scan(&answer)

	questionAnswer[question] = answer
}
