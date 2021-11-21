package main

import (
	"log"
	"net/http"
	"os"
)

func RouteStart() {

	if os.Getenv("ENV_MODE") == "dev" {
		log.Println("ENV development implemented!")
	}

	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	ServeWs(w, r)
	// })

	// app.Listen(":" + os.Getenv("APP_RUN_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("APP_RUN_PORT"), nil))
	log.Println("Service started!")

}
