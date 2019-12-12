package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func ErrorMessage(message string) map[string]interface{} {
	return map[string]interface{}{"result": "error", "message": message}
}

func Respond(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func MatchString(pattern, s string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(s)
}
