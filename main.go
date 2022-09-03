package main

import (
	"eaza-go/database"
	"eaza-go/internal/common"
	"eaza-go/internal/course"
	"github.com/bytedance/sonic"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	_ = godotenv.Load()

	app := fiber.New(fiber.Config{
		JSONDecoder: sonic.Unmarshal,
		JSONEncoder: sonic.Marshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewErrorResponse(code, err))
			}

			// Return from handler
			return nil
		},
	})

	app.Use(cache.New(cache.Config{
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.Request().URI().String()
		},
		Storage: database.NewRedisClient(&redis.Options{
			Addr: os.Getenv("REDIS_URL"),
		}),
		Expiration: time.Minute,
	}))

	database.Connect()

	v1 := app.Group("/v1")

	registerPlugin(v1)

	log.Fatal(app.Listen(":3000"))
}

func registerPlugin(r fiber.Router) {
	g := r.Group("/plugin")
	c := course.NewController(&course.ServiceImpl{DB: database.DB})
	g.Get("/course", c.GetCourseByAbbrAndNumber)
	g.Get("/grade", c.GetGradesBySubject)
}
