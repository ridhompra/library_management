package router

import (
	"log"
	"net/http"
	"project/library_Management/controller"
	"project/library_Management/controller/productcontroller"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	port := ":8080"
	r.HandleFunc("/", controller.Hompage)

	r.HandleFunc("/book", productcontroller.GetAllBook).Methods("GET")
	r.HandleFunc("/book/{id}", productcontroller.GetBookbyid).Methods("GET")
	r.HandleFunc("/book", productcontroller.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", productcontroller.UpdateBook).Methods("PUT")
	r.HandleFunc("/book", productcontroller.DeleteBook).Methods("DELETE")

	r.HandleFunc("/employee", productcontroller.GetAllEmployee).Methods("GET")
	r.HandleFunc("/employee/{id}", productcontroller.GetEmployeebyid).Methods("GET")
	r.HandleFunc("/employee", productcontroller.CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{id}", productcontroller.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee", productcontroller.DeleteEmployee).Methods("DELETE")

	r.HandleFunc("/visitor", productcontroller.GetAllVisitor).Methods("GET")
	r.HandleFunc("/visitor/{id}", productcontroller.GetAllVisitor).Methods("GET")
	r.HandleFunc("/visitor", productcontroller.CreateVisitor).Methods("POST")
	r.HandleFunc("/visitor/{id}", productcontroller.UpdateVisitor).Methods("PUT")
	r.HandleFunc("/visitor", productcontroller.DeleteVisitor).Methods("DELETE")

	log.Printf("Server Running port %s \n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
