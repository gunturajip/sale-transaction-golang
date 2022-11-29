package main

import (
	"fmt"
	"log"
	"tugas_akhir/internal/infrastructure/container"

	rest "tugas_akhir/internal/server/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fmt.Println("Hello")

	containerConf := container.InitContainer()
	app := fiber.New()
	app.Use(logger.New())

	rest.HTTPRouteInit(app, containerConf)

	log.Fatal(app.Listen(":8000"))
}
