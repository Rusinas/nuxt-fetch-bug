package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Append("Access-Control-Allow-Origin", "*")
		return c.SendStatus(fiber.StatusOK)
	})


	log.Fatal(app.Listen(":4000"))
}
