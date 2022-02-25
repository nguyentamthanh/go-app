package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tamthanh/go-app/database"
	"github.com/tamthanh/go-app/router"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("hello thanh page")
}
func setupRouters(app *fiber.App) {
	//welcome
	app.Get("/api", welcome)
	//book
	app.Post("/api/products", router.CreateProduct)
	app.Get("/api/products", router.Getproducts)
	app.Get("/api/products/:id", router.GetProduct)
	app.Delete("/api/products:id", router.DeleteProduct)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRouters(app)
	log.Fatal(app.Listen(":7000"))
}
