package users

import (
	"fmt"
	"shared-fridge-be/pkg/router"
)

func RegisterRoutes(prefix string, handler *UserHandler, router *router.Router) {
	router.Get(fmt.Sprintf("/%s", prefix), handler.GetUsers)
}
