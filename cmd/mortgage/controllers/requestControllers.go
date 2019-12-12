package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mortgage/cmd/mortgage/common"
	"mortgage/cmd/mortgage/models"
	u "mortgage/cmd/mortgage/utils"
	"net/http"
)

var CreateRequest = func(w http.ResponseWriter, r *http.Request) {

	request := &common.Request{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.ErrorMessage("Error request decode"))
	}

	resp := request.Create()
	u.Respond(w, http.StatusCreated, resp)

}

var GetStatusByRequestId = func(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	mr := models.GetStatusByRequestId(id)
	if mr == nil {
		u.Respond(w, http.StatusNotFound, u.ErrorMessage("request not found"))
	}

	u.Respond(w, http.StatusOK, mr)

}

var GetRequests = func(w http.ResponseWriter, r *http.Request) {

	requests := models.GetRequests()
	if requests == nil {
		u.Respond(w, http.StatusNotFound, u.ErrorMessage("error requests founding"))
	}

	u.Respond(w, http.StatusOK, requests)

}
