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

	pb "server/protos"

	"os"
	"strconv"
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

func (s *Server) RunServer() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	conn, err := grpc.Dial(
		os.Getenv("GRPC_LOGGER_HOST"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	LoggerClient := pb.NewLoggerClient(conn)

	app.Get("/", func(c *fiber.Ctx) error {

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

	err = app.Listen(fmt.Sprint(s.addr, ":", s.port))
	if err != nil {
		log.Fatalf("Error on opening port")
	}
}

func main() {
	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Error on parsing port")
	}

	s := newServer("0.0.0.0", PORT)
	s.RunServer()
}
