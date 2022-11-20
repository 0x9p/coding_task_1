package util

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	JsonResponse(w, r, code, map[string]string{"error": err.Error()})
}

func PartialError(w http.ResponseWriter, r *http.Request, code int, err error, data interface{}) {
	JsonResponse(w, r, code, map[string]interface{}{"error": err.Error(), "data": data})
}

func Response(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func JsonResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	Response(w, code)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		Error(w, r, http.StatusInternalServerError, err)
	}
}
