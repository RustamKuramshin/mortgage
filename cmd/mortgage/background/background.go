package background

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mortgage/cmd/mortgage/common"
	c "mortgage/cmd/mortgage/concurrent"
	"net/http"
	"time"
)

func sendMortgageRequest() {
	for range time.Tick(time.Duration(1) * time.Second) {
		if !c.RequestsQueue.IsEmpty() {

			request := c.RequestsQueue.Peek()

			reqBody, err := json.Marshal(request)
			if err != nil {
				continue
			}
			resp, err := http.Post("http://localhost:9000/request", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
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
	for range time.Tick(time.Duration(1) * time.Second) {
		if !c.StatusesQueue.IsEmpty() {
			requestId := c.StatusesQueue.Peek()

			resp, err := http.Get(fmt.Sprintf("http://localhost:9000/request/%s", requestId))
			if err != nil {
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
