package main

import (
	"log"

	api "github.com/Killerrekt/User_Crud_in_go/API"
	database "github.com/Killerrekt/User_Crud_in_go/Database"
	"github.com/gofiber/fiber/v2"
)

func Setuproutes(app *fiber.App) {
	app.Post("/create", api.Createuser)
	app.Put("/update", api.Updateuser)
	app.Delete("/delete", api.Deleteuser)
	app.Get("/user", api.Getuser)
	app.Get("/check", api.Checkuser)
}

func main() {
	app := fiber.New()

	err := database.InitDatabase()

	if err != nil {
		log.Fatal("Something went wrong while connecting to DB")
	}

	Setuproutes(app)
	app.Listen(":3000")
}
