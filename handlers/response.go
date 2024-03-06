package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonResponse map[string]interface{}

func WriteJSON(w http.ResponseWriter, r *http.Request, message string, statusCode int, data interface{}, count interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := jsonResponse{
		"message": message,
		"status":  statusCode,
		"data":    data,
		"count":   count,
		"links":   hateoasLinks(r),
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	response := jsonResponse{
		"error":  message,
		"status": statusCode,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func hateoasLinks(r *http.Request) jsonResponse {
	return jsonResponse{
		"rel":  "tasks",
		"path": r.RequestURI,
		"type": r.Method,
	}
}
