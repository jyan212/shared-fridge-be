package router

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	Handler *http.ServeMux
	routes  map[string]string
}

type Handler func(http.ResponseWriter, *http.Request)

func NewRouter() *Router {
	routes := make(map[string]string)

	return &Router{
		Handler: http.NewServeMux(),
		routes:  routes,
	}
}

func (router *Router) Get(path string, handler Handler) {
	router.routes[path] = "GET"
	router.Handler.HandleFunc(fmt.Sprintf("GET %s", path), handler)
}

func (router *Router) Post(path string, handler Handler) {
	router.routes[path] = "POST"
	router.Handler.HandleFunc(fmt.Sprintf("POST %s", path), handler)
}

func (router *Router) Patch(path string, handler Handler) {
	router.routes[path] = "PATCH"
	router.Handler.HandleFunc(fmt.Sprintf("PATCH %s", path), handler)
}

func (router *Router) Delete(path string, handler Handler) {
	router.routes[path] = "DELETE"
	router.Handler.HandleFunc(fmt.Sprintf("DELETE %s", path), handler)
}

func (router *Router) PrintRoutes() {
	log.Println("********* Registered Routes **************")
	for path, method := range router.routes {
		log.Printf("%s - %s", method, path)
	}
	log.Println("******************************************")
}
