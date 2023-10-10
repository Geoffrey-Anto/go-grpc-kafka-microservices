package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/geoffrey-anto/golang-microservice-apis/protos"
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
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		conn, err := grpc.Dial(
			"localhost:50051",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		LoggerClient := pb.NewLoggerClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := LoggerClient.SaveLog(ctx, &pb.LogSaveRequest{
			Id:   c.IP(),
			Time: time.Now().String(),
			Log:  "GET /",
		})
		if err != nil {
			log.Fatalf("could not log: %v", err)
		}
		log.Printf("Log Success: %v", r.Success)

		return c.Render("index", fiber.Map{})
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
