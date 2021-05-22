package util

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, data interface{}, statusCode int) {
	/* set map response */
	response, _ := json.Marshal(map[string]interface{}{"data": data})
	/* response json */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func ResponseError(w http.ResponseWriter, errorMessage string, statusCode int) {
	/* set map response */
	response, _ := json.Marshal(map[string]string{"message": errorMessage})
	/* response json */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
