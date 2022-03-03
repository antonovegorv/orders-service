package main

import (
	"encoding/json"
	"log"

	"github.com/antonovegorv/orders-service/cache"
	"github.com/antonovegorv/orders-service/config"
	"github.com/antonovegorv/orders-service/database"
	"github.com/antonovegorv/orders-service/internal/model"
	"github.com/antonovegorv/orders-service/router"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/stan.go"
)

func main() {
	// Connect to the Database.
	if err := database.Connect(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Connected to the Database!")
	}

	// Initialize the cache for the service to work.
	if err := cache.Init(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Cache has been initialized!")
	}

	// Connect to the nats-streaming-server.
	sc, err := stan.Connect(config.Get("NATS_CLUSTER_ID"), config.Get("NATS_CLIENT_ID"))
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	// Subscribe to the channel in order to start listening to messages.
	sc.Subscribe(config.Get("NATS_SUBJECT"), messageHandler, stan.DurableName("processed"))

	// Create an instance of the fiber application.
	app := fiber.New()

	// Setup basic handler for a homepage route.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🦆 Quack-quack...")
	})

	// Setup routes for the entire app.
	router.SetupRoutes(app)

	// Start the HTTP server to listen for incoming requests.
	log.Fatal(app.Listen(config.Get("HTTP_PORT")))
}

// A handler for an incoming message from nats-streaming-server.
func messageHandler(m *stan.Msg) {
	// Create a variable to store order in.
	var order model.Order

	// Try to unmarshal data received by the nats.
	if err := json.Unmarshal(m.Data, &order); err != nil {
		log.Println("Unpredicted data has been received. Skipped.")
		return
	}

	// Write new value to the Database.
	_, err := database.DB.Exec("INSERT INTO orders (data) VALUES($1)", order)
	if err != nil {
		// Log that we have an error in inserting data to the Database.
		log.Println(err) // Probably should to be Fatalln...
	} else {
		// Log that we have inserted data to the Database.
		log.Printf("New order (%v) has been inserted to the Database!\n", order.UUID)
	}

	// Set new value to cache.
	if err := cache.Orders.Set(order.UUID.String(), order); err != nil {
		// Log that we have an error in caching data.
		log.Println(err) // Probably should to be Fatalln...
	} else {
		// Log that we have set new value to cache.
		log.Printf("New order (%v) has been successfully cached!\n", order.UUID)
	}
}
