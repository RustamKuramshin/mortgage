package models

import (
	"github.com/jinzhu/gorm"
	u "mortgage/cmd/mortgage/utils"
	"regexp"
	"strings"
)

type MortgageRequest struct {
	Request *Request `json:"request"`
}

type Request struct {
	gorm.Model
	Id         string `json:"omitempty"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	MiddleName string `json:"middlename"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

type MortgageRequestResponse struct {
	RequestResponse *RequestResponse `json:"request"`
}

type RequestResponse struct {
	Id         string `json:"id"`
	StatusCode string `json:"status_code"`
}

func (r *Request) Validate() (map[string]interface{}, bool) {

	if strings.TrimSpace(r.FirstName) == "" {
		return u.ErrorMessage("first name must be specified"), false
	}
	if strings.TrimSpace(r.LastName) == "" {
		return u.ErrorMessage("last name must be specified"), false
	}
	if strings.TrimSpace(r.MiddleName) == "" {
		return u.ErrorMessage("middle name must be specified"), false
	}

	if phone := strings.TrimSpace(r.Phone); phone == "" {
		return u.ErrorMessage("phone must be specified"), false
	} else {
		if ok, _ := regexp.MatchString(`\+79\d{9}`, phone); !ok {
			return u.ErrorMessage("phone number must be in international format"), false
		}
	}

	if email := strings.TrimSpace(r.Email); email == "" {
		return u.ErrorMessage("email must be specified"), false
	} else {
		if ok, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email); !ok {
			return u.ErrorMessage("email must be correct"), false
		}
	}

	return u.ErrorMessage(""), true
}

func (r *Request) Create() interface{} {

	if resp, ok := r.Validate(); !ok {
		return resp
	}

	GetDB().Create(r)

	return RequestResponse{Id: r.ID, StatusCode: ""}
}