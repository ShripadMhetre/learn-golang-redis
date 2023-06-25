package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"golang-redis/middlewares"
	"golang-redis/redis"
	"golang-redis/utils"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.TODO()
	redis.ConnectRedis(ctx)
	app := fiber.New()

	app.Get("/:id", middlewares.VerifyCache, func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := http.Get("https://jsonplaceholder.typicode.com/users/" + id)
		if err != nil {
			return err
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cacheErr := redis.RedisClient.Set(ctx, id, body, 100*time.Second).Err()
		if cacheErr != nil {
			return cacheErr
		}

		data := utils.ToJson(body)
		return c.JSON(fiber.Map{"Data": data})
	})

	app.Listen(":3000")
}
