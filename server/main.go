package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	addr string
	port int
}

func NewServer(addr string, port int) *Server {
	return &Server{
		port: port,
		addr: addr,
	}
}

func (s *Server) RunServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		c.SendString("OK")
		return c.SendStatus(200)
	})

	err := app.Listen(fmt.Sprint(s.addr, ":", s.port))
	if err != nil {
		log.Fatalf("Error on opening port")
	}
}
