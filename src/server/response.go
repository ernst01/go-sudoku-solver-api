package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type appError struct {
	Type    string `json:"type"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	HelpURL string `json:"help_url"`
}

func sendSuccess(w http.ResponseWriter, httpStatus int, data interface{}) {
	var buffer bytes.Buffer

	if data != nil {
		if err := json.NewEncoder(&buffer).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	if _, err := io.Copy(w, &buffer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sendError(w http.ResponseWriter, httpStatus int, format string, a ...interface{}) {
	errorObj := appError{
		Type:    "error",
		Status:  httpStatus,
		Message: fmt.Sprintf(format, a...),
		Code:    slugify(http.StatusText(httpStatus)),
		HelpURL: "https://www.google.com",
	}
	jsonError, err := json.Marshal(errorObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Error(w, string(jsonError), httpStatus)
}

func slugify(message string) string {
	message = strings.ToLower(message)
	message = strings.Replace(message, " ", "_", -1)
	return message
}
