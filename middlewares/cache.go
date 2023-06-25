package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"golang-redis/redis"
	"golang-redis/utils"
)

func VerifyCache(c *fiber.Ctx) error {
	id := c.Params("id")
	val, err := redis.Cache.Get(redis.Ctx, id).Bytes()
	if err != nil {
		return c.Next()
	}

	data := utils.ToJson(val)
	return c.JSON(fiber.Map{"Cached": data})
}
