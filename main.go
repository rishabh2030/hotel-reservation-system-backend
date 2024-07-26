package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rishabh2030/hotel-reservation-system/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The Listener of The API Server to listen")
	flag.Parse()
	app := fiber.New()

	apiV1 := app.Group("/api/v1/")

	apiV1.Get("/users", api.HandleListGetUser)
	apiV1.Get("/user/:id", api.GetUserHandler)
	log.Fatal(app.Listen(*listenAddr))
}
