package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	erpCatCodeTopicMap := map[int]string{
		11:   "Academic",
		12:   "Administrative",
		13:   "Miscellaneous",
		1001: "Academic (UG) section notices",
		1002: "Academic (PG) section notices",
	}

	endpoint := "https://erp.iitkgp.ac.in/InfoCellDetails/internal_noticeboard/get_notice_list.htm?cat_code=%d"

	resp, err := http.Get(fmt.Sprintf(endpoint, ""))

	if err != nil {
		log.Printf(err.Error())
	}

	fmt.Println(resp.StatusCode)

}
