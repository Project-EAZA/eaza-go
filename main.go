package main

import (
	"eaza-go/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	_ = godotenv.Load()

	app := fiber.New()

	database.Connect()

	log.Fatal(app.Listen(":3000"))
}
