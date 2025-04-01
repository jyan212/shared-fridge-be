package app

import (
	"fmt"
	"shared-fridge-be/pkg/router"
)

func RegisterRoutes(prefix string, handler *AppHandler, router *router.Router) {
	router.Get(fmt.Sprintf("/%s/healthcheck", prefix), handler.HealthCheck)
}
