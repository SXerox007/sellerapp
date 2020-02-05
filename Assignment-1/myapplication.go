package main

import (
	"log"
	"sellerapp/base/db/mongodb"
	env "sellerapp/base/environment"
	"sellerapp/base/router"
	"sellerapp/base/router/server"
)

// init
func Init() {
	environment := env.GetEnv()
	port := env.GetPort()
	mongodbSetup()
	setupRouter(environment, port)

}

func main() {
	Init()
}

// mongodb setup
func mongodbSetup() {
	if err := mongodb.InitDB(); err != nil {
		log.Println("Error in Init MongoDB:", err)
		return
	}
}

func setupRouter(env, port string) {
	//initilize the router
	router.InitRouter()
	sellerMux := router.SubRouter("/sellerapp")
	sellerMux.HandleFunc("/{version}/order", OrderDataPost()).Methods("POST")
	log.Println("Server serve at", env+":"+port)
	server.StartServer(port)
}
