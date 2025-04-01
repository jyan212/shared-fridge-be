package app

import (
	"log"
	"shared-fridge-be/pkg/router"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// handler -> service -> repo -> model -> db
type AppModule struct {
	handler *AppHandler
}

func NewModule() *AppModule {
	return &AppModule{
		handler: &AppHandler{},
	}
}

func (app AppModule) SetupModule(router *router.Router, dbClient *mongo.Client) error {
	log.Println("Setting up App Module")
	// Dependency Injections
	appHandler := SetupHandler()

	RegisterRoutes("app", appHandler, router)

	return nil
}
