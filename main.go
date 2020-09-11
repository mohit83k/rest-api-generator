package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohit83k/rest-api-generator/server"

	. "github.com/mohit83k/rest-api-generator/configuration"
)

type apiRouter struct {
	fileName string
}

func (a apiRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request : ", r.RequestURI)
	fmt.Println("Servering : ", a.fileName)
	http.ServeFile(rw, r, a.fileName)
}

func main() {

	server.AddRouter(addRouter)

	server.Start()
}

func addRouter(router *mux.Router) {
	for _, api := range Config.ApiRoutes {
		router.Handle(api.Url, apiRouter{api.File})
	}

	router.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
	})
}
