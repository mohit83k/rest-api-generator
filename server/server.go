package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/mohit83k/rest-api-generator/configuration"
)

var routers []func(*mux.Router)

//AddRouter to the Server
func AddRouter(r func(*mux.Router)) {
	routers = append(routers, r)
}

//Start the server
func Start() {
	r := mux.NewRouter()
	for _, router := range routers {
		router(r)
	}

	allowedHeader := handlers.AllowedHeaders(Config.Server.AllowedHeaders)
	allowedOrigin := handlers.AllowedOrigins(Config.Server.CorsException)

	server := &http.Server{
		Addr:    ":" + Config.Server.Port,
		Handler: handlers.CORS(allowedHeader, allowedOrigin)(r),
	}

	var err error
	log.Println("Starting on Port : ", Config.Server.Port)
	if Config.Server.Tls {
		log.Printf("TLS Enabled")
		err = server.ListenAndServeTLS(Config.Server.Certificate, Config.Server.PrivateKey)
	} else {
		err = server.ListenAndServe()
	}

	if err != nil {
		log.Println(err)
	}

}
