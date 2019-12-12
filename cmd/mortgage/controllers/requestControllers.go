package controllers

import (
	"encoding/json"
	"mortgage/cmd/mortgage/models"
	u "mortgage/cmd/mortgage/utils"
	"net/http"
)

var CreateRequest = func(w http.ResponseWriter, r *http.Request) {

	request := &models.Request{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.ErrorMessage("Error request decode"))
	}

	resp := request.Create()
	u.Respond(w, http.StatusCreated, resp)

}
