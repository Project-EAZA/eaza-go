package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	v1 := app.Group("/v1")

	course := v1.Group("/course")
	course.Get("")
}
