package users

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userService *UserService
}

func (userHandler *UserHandler) SetupHandler(userService *UserService) *UserHandler {
	userHandler.userService = userService
	return userHandler
}

func (UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
		"users":  "Jun Yan",
	})
}
