package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/antonovegorv/orders-service/cache"
	"github.com/antonovegorv/orders-service/config"
	"github.com/antonovegorv/orders-service/database"
	"github.com/antonovegorv/orders-service/internal/model"
	"github.com/antonovegorv/orders-service/router"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/stan.go"
)

// Create 3-level logging system.
var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// Connect to the Database.
	if err := database.Connect(); err != nil {
		ErrorLogger.Fatalln(err)
	} else {
		InfoLogger.Println("Connected to the Database!")
	}

	// Initialize the cache for the service to work.
	if err := cache.Init(); err != nil {
		ErrorLogger.Fatalln(err)
	} else {
		InfoLogger.Println("Cache has been initialized!")
	}

	// Connect to the nats-streaming-server.
	sc, err := stan.Connect(config.Get("NATS_CLUSTER_ID"), config.Get("NATS_CLIENT_ID"))
	if err != nil {
		InfoLogger.Fatalln(err)
	}
	defer sc.Close()

	// Subscribe to the channel in order to start listening to messages.
	sc.Subscribe(config.Get("NATS_SUBJECT"), messageHandler, stan.DurableName("processed"))

	// Create an instance of the fiber application.
	app := fiber.New()

	// Setup serving static files.
	app.Static("/", "./public")

	// Setup routes for the entire app.
	router.SetupRoutes(app)

	// Start the HTTP server to listen for incoming requests.
	ErrorLogger.Fatalln(app.Listen(config.Get("HTTP_PORT")))
}

// A handler for an incoming message from nats-streaming-server.
func messageHandler(m *stan.Msg) {
	// Create a variable to store order in.
	var order model.Order

	// Try to unmarshal data received by the nats.
	//
	// So, this check will only protect us from some inappropriate information.
	// However, we do not have protection on JSON of a different format.
	// To do this, it is worth implementing either a check of all fields for
	// compliance. Or initially specify field references in order to check the
	// value for nil.
	if err := json.Unmarshal(m.Data, &order); err != nil {
		WarningLogger.Println("Unpredicted data has been received. Skipped.")
		return
	}

	// Write new value to the Database.
	_, err := database.DB.Exec("INSERT INTO orders (data) VALUES($1)", order)
	if err != nil {
		// Log that we have an error in inserting data to the Database.
		WarningLogger.Println(err) // Probably should to be Fatalln...
	} else {
		// Log that we have inserted data to the Database.
		InfoLogger.Printf("New order (%v) has been inserted to the Database!\n", order.UUID)
	}

	// Set new value to cache.
	if err := cache.Orders.Set(order.UUID.String(), order); err != nil {
		// Log that we have an error in caching data.
		WarningLogger.Println(err) // Probably should to be Fatalln...
	} else {
		// Log that we have set new value to cache.
		InfoLogger.Printf("New order (%v) has been successfully cached!\n", order.UUID)
	}
}
