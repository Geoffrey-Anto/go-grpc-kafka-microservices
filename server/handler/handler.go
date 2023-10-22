package handler

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	// pbLogger "server/protos/logger"
	pbRandomJoke "server/protos/randomjoke"
)

func MainHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func HealthHandler(c *fiber.Ctx) error {
	c.SendString("OK")
	return c.SendStatus(200)
}

func RandomJokeHandler(c *fiber.Ctx, RandomJokeClient pbRandomJoke.RandomJokeServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := RandomJokeClient.GetRandomJoke(ctx, &pbRandomJoke.RandomJokeRequest{
		Category: "dev",
		Timeout:  5,
	})
	if err != nil {
		log.Fatalf("could not get random joke: %v", err)
	}

	c.SendString(r.Joke)
	return c.SendStatus(200)
}
