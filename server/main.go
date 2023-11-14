package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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

func (s *Server) RunServer(LoggerClient pbLogger.LoggerClient, RandomJokeClient pbRandomJoke.RandomJokeServiceClient, producer *kafka.Producer) {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(func(c *fiber.Ctx) error {
		// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		// defer cancel()
		// ip_addr := c.Request().Header.Peek("X-Forwarded-For")
		// if ip_addr == nil {
		// 	ip_addr = []byte("unknown")
		// }
		// _, err := LoggerClient.SaveLog(ctx, &pbLogger.LogSaveRequest{
		// 	Id:   string(ip_addr[:]),
		// 	Time: time.Now().String(),
		// 	Log:  c.Path(),
		// })
		// if err != nil {
		// 	log.Fatalf("could not log: %v", err)
		// }

		kafkaTopic := os.Getenv("KAFKA_TOPIC")

		ip_addr := c.Request().Header.Peek("X-Forwarded-For")
		if ip_addr == nil {
			ip_addr = []byte("unknown")
		}

		ip_addr_str := string(ip_addr[:])

		value := fmt.Sprintf(
			"%+v, %+v, %+v at %+v from %+v\n",
			c.Hostname(),
			c.Path(),
			c.Method(),
			time.Now().String(),
			ip_addr_str,
		)

		producerErr := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
			Value:          []byte(value),
		}, nil)

		if producerErr != nil {
			fmt.Println("unable to enqueue message")
		}
		event := <-producer.Events()

		message := event.(*kafka.Message)

		if message.TopicPartition.Error != nil {
			fmt.Println("Delivery failed due to error ", message.TopicPartition.Error)
		} else {
			fmt.Println("Delivered message to offset " + message.TopicPartition.Offset.String() + " in partition " + message.TopicPartition.String())
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

	app.Get("/host", func(c *fiber.Ctx) error {
		return handler.GetHost(c)
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

	kafkaServer := os.Getenv("KAFKA_SERVER")

	producer, producerCreateErr := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})

	if producerCreateErr != nil {
		fmt.Println("Failed to create producer due to ", producerCreateErr)
		os.Exit(1)
	}

	s := newServer("0.0.0.0", PORT)
	s.RunServer(LoggerClient, RandomJokeClient, producer)
}
