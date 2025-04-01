package container

import (
	"log"
	"net/http"
	"shared-fridge-be/pkg/router"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AppContainer struct {
	Modules []AppModule
}

func NewApp() *AppContainer {
	return &AppContainer{}
}

func (ac *AppContainer) Add(module AppModule) {
	ac.Modules = append(ac.Modules, module)
}

func (ac *AppContainer) RegisterModules(modules []AppModule) {
	ac.Modules = append(ac.Modules, modules...)
}

func (ac AppContainer) StartApp(router *router.Router, dbClient *mongo.Client, addr string) {
	log.Println("Starting App..")
	for _, module := range ac.Modules {
		if err := module.SetupModule(router, dbClient); err != nil {
			log.Fatalln("Failed to start the App. ", err)
		}
	}

	router.PrintRoutes()

	log.Println("Started the App and listening on ", addr)
	http.ListenAndServe("localhost:8080", router.Handler)

}
