package users

import (
	"log"
	"shared-fridge-be/pkg/router"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// handler -> service -> repo -> model -> db
type UserModule struct {
	handler *UserHandler
	service *UserService
	repo    *UserRepo
	model   *UserModel
}

func NewModule() *UserModule {
	return &UserModule{
		handler: &UserHandler{},
		service: &UserService{},
		repo:    &UserRepo{},
		model:   &UserModel{},
	}
}

func (user UserModule) SetupModule(router *router.Router, dbClient *mongo.Client) error {
	log.Println("Setting up User Module")
	// Dependency Injections
	userModel := user.model.SetupModel(dbClient)
	userRepo := user.repo.SetupRepo(userModel)
	userService := user.service.SetupService(userRepo)
	userHandler := user.handler.SetupHandler(userService)

	RegisterRoutes("users", userHandler, router)

	return nil
}
