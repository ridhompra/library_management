package router

import (
	"log"
	"net/http"
	"project/library_Management/controller"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	port := ":8080"
	r.HandleFunc("/", controller.Hompage)

	log.Printf("Server Running port %s \n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
