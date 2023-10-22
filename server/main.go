package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbLogger "server/protos/logger"
	pbRandomJoke "server/protos/randomjoke"

	"os"
	"strconv"

	"server/handler"
)

type Server struct {
	addr string
	port int
}

func newServer(addr string, port int) *Server {
	return &Server{
		port: port,
		addr: addr,
	}
}

func (s *Server) RunServer(LoggerClient pbLogger.LoggerClient, RandomJokeClient pbRandomJoke.RandomJokeServiceClient) {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err := LoggerClient.SaveLog(ctx, &pbLogger.LogSaveRequest{
			Id:   c.IP(),
			Time: time.Now().String(),
			Log:  c.Path(),
		})
		if err != nil {
			log.Fatalf("could not log: %v", err)
		}
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return handler.MainHandler(c)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return handler.HealthHandler(c)
	})

	app.Post("/random-joke", func(c *fiber.Ctx) error {
		return handler.RandomJokeHandler(c, RandomJokeClient)
	})

	err := app.Listen(fmt.Sprint(s.addr, ":", s.port))
	if err != nil {
		log.Fatalf("Error on opening port")
	}
}

func main() {
	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Error on parsing port")
	}

	loggerConn, err := grpc.Dial(
		os.Getenv("GRPC_LOGGER_HOST"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer loggerConn.Close()
	LoggerClient := pbLogger.NewLoggerClient(loggerConn)

	randomJokeConn, err := grpc.Dial(
		os.Getenv("GRPC_RANDOMJOKE_HOST"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer randomJokeConn.Close()

	RandomJokeClient := pbRandomJoke.NewRandomJokeServiceClient(randomJokeConn)

	s := newServer("0.0.0.0", PORT)
	s.RunServer(LoggerClient, RandomJokeClient)
}
