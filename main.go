package main

import (
	"log"
	appContainer "shared-fridge-be/internal/app"
	"shared-fridge-be/internal/users"
	"shared-fridge-be/pkg/container"
	"shared-fridge-be/pkg/db/mongodb"
	"shared-fridge-be/pkg/router"
)

func main() {
	if conn, err := mongodb.NewConnection("mongodb://localhost:3366"); err != nil {
		log.Fatalln("Failed to start server")
	} else {
		router := router.NewRouter()
		app := container.NewApp()

		// replace this code to get modules from config
		app.RegisterModules([]container.AppModule{
			appContainer.NewModule(),
			users.NewModule(),
		})
		app.StartApp(router, conn.Client, "localhost:8080")
	}
}
