package main

import (
	"project/library_Management/models"
	"project/library_Management/router"
)

func main() {
	models.ConnectionDB()
	router.Router()
}
