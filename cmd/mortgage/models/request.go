package models

import (
	"fmt"
	c "mortgage/cmd/mortgage/common"
)

type MortgageRequest struct {
	Request *c.Request `json:"request"`
}

type MortgageRequestResponse struct {
	RequestResponse *c.RequestResponse `json:"request"`
}

func GetStatusByRequestId(id string) *c.RequestResponse {

	request := new(c.Request)
	err := c.GetDB().Table("requests").Where("bank_id_request = ?", id).Find(request).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &c.RequestResponse{Id: request.ID.String(), StatusCode: request.StatusCode}
}

func GetRequests() []*c.Request {

	requests := make([]*c.Request, 0)

	err := c.GetDB().Table("requests").Find(&requests).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return requests
}
