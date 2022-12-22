package router

import (
	"log"
	"os"
	"project/library_Management/controller/authcontrollers"
	"project/library_Management/controller/productcontrollers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Router() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect .ENV")
	}

	port := os.Getenv("PORT")
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// v1
	// sign up and login session
	v1.Get("/", authcontrollers.HealthCheck)
	v1.Post("/signup", authcontrollers.SignUp)
	v1.Get("/users", authcontrollers.GetAllDataUser)
	v1.Post("/login", authcontrollers.Login)

	//book
	v1.Get("/book", productcontrollers.GetAllBook)
	v1.Get("/book/:id", productcontrollers.GetBookbyid)
	v1.Post("/book", productcontrollers.CreateBook)
	v1.Put("/book/:id", productcontrollers.UpdateBook)
	v1.Delete("/book", productcontrollers.DeleteBook)

	// employee
	v1.Get("/employee", productcontrollers.GetAllEmployee)
	v1.Get("/employee/:id", productcontrollers.GetEmployeebyid)
	v1.Post("/employee", productcontrollers.CreateEmployee)
	v1.Put("/employee/:id", productcontrollers.UpdateEmployee)
	v1.Delete("/employee", productcontrollers.DeleteEmployee)

	// visitor
	v1.Get("/visitor", productcontrollers.GetAllVisitor)
	v1.Get("/visitor/:id", productcontrollers.GetVisitorById)
	v1.Post("/visitor", productcontrollers.CreateVisitor)
	v1.Put("/visitor/:id", productcontrollers.UpdateVisitor)
	v1.Delete("/visitor", productcontrollers.DeleteVisitor)

	log.Printf("Server Running port %s \n", port)
	log.Println(app.Listen(port))

}
