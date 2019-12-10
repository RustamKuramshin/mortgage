package models

import "github.com/jinzhu/gorm"

type MortgageRequest struct {
	Request *Request `json:"request"`
}

type Request struct {
	gorm.Model
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	MiddleName string `json:"middlename"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

func (r *Request) Vaildate() {

}
