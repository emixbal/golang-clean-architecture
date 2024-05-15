package main

import (
	"golang-clean-architecture/internal/component"
	"golang-clean-architecture/internal/config"
	"golang-clean-architecture/internal/module/customer"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.Get()
	dbConnection := component.GetDbConnection(conf)

	app := fiber.New()

	customerRepository := customer.NewRepository(dbConnection)
	customerService := customer.NewService(customerRepository)
	customer.NewApi(app, customerService)

	app.Listen(":3000")
}
