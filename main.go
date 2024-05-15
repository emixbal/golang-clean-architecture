package main

import (
	"golang-clean-architecture/internal/component"
	"golang-clean-architecture/internal/config"
	"golang-clean-architecture/internal/module/customer"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	conf := config.Get()
	dbConnection := component.GetDbConnection(conf)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${locals:requestid}] ${ip} - ${method} ${status} ${path} ${latency} \n",
	}))

	customerRepository := customer.NewRepository(dbConnection)
	customerService := customer.NewService(customerRepository)
	customer.NewApi(app, customerService)

	app.Listen(":3000")
}
