package app

import (
	"encoding/json"
	"net/http"
)

type AppHandler struct{}

func SetupHandler() *AppHandler {
	return &AppHandler{}
}

func (AppHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}
