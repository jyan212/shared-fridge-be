package container

import (
	"shared-fridge-be/pkg/router"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AppModule interface {
	SetupModule(router *router.Router, dbClient *mongo.Client) error
}
