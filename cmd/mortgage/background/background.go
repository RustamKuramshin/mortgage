package background

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mortgage/cmd/mortgage/common"
	c "mortgage/cmd/mortgage/concurrent"
	"net/http"
	"os"
	"strconv"
	"time"
)

func sendMortgageRequest() {
	rps, _ := strconv.Atoi(os.Getenv("BANK_SEND_REQUEST_RPS"))
	for range time.Tick(time.Duration(rps) * time.Second) {
		if !c.RequestsQueue.IsEmpty() {

			request := c.RequestsQueue.Peek()

			reqBody, err := json.Marshal(request)
			if err != nil {
				continue
			}

			url := os.Getenv("BANK_SEND_REQUEST_URL")
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				log.Println(fmt.Sprintf("POST %s", url))
				log.Println("bank api not available")
				continue
			}

			defer resp.Body.Close()

			rr := new(common.RequestResponse)
			err = json.NewDecoder(resp.Body).Decode(rr)
			if err != nil {
				continue
			}

			if rr.StatusCode == "processing" {
				r := request.(common.Request)
				common.GetDB().Model(&r).Update("bank_id_request", rr.Id)
				common.GetDB().Model(&r).Update("status_code", "processing")
				c.RequestsQueue.Dequeue()
				c.StatusesQueue.Enqueue(rr.Id)
			}

		}
	}
}

func checkStatusesRequests() {
	rps, _ := strconv.Atoi(os.Getenv("BANK_CHECK_STATUS_RPS"))
	for range time.Tick(time.Duration(rps) * time.Second) {
		if !c.StatusesQueue.IsEmpty() {
			requestId := c.StatusesQueue.Peek()

			url := fmt.Sprintf("%s/%s", os.Getenv("BANK_CHECK_STATUS_URL"), requestId)
			resp, err := http.Get(url)
			if err != nil {
				log.Println(fmt.Sprintf("GET %s", url))
				log.Println("bank api not available")
				continue
			}

			defer resp.Body.Close()

			rr := new(common.RequestResponse)
			err = json.NewDecoder(resp.Body).Decode(rr)
			if err != nil {
				continue
			}

			var r common.Request
			common.GetDB().First(&r, "bank_id_request = ?", rr.Id)
			common.GetDB().Model(&r).Update("status_code", rr.StatusCode)
			c.StatusesQueue.Dequeue()
		}
	}
}

func StartBackgroundTasks() {
	go sendMortgageRequest()
	go checkStatusesRequests()
}
