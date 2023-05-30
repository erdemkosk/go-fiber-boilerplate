package main

import (
	"fmt"
	"go-fiber-boilerplate/api"
	"go-fiber-boilerplate/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Server initialization
	app := server.Create()

	// Migrations
	//database.DB.AutoMigrate(&books.Book{})

	// Api routes
	api.Setup(app)

	go func() {
		if err := server.Listen(app); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
