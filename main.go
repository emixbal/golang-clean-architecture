package main

import (
	"golang-clean-architecture/internal/component"
	"golang-clean-architecture/internal/config"
	"golang-clean-architecture/internal/module/customer"
	"golang-clean-architecture/internal/module/history"
	"golang-clean-architecture/internal/module/vehicle"

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
		Format: "[${locals:requestid}] ${ip} - ${method} ${status} ${latency} ${path}\n",
	}))

	customerRepository := customer.NewRepository(dbConnection)
	historyRepository := history.NewRepository(dbConnection)
	vehicleRepository := vehicle.NewRepository(dbConnection)

	customerService := customer.NewService(customerRepository)
	vehicleService := vehicle.NewService(vehicleRepository, historyRepository)

	customer.NewApi(app, customerService)
	vehicle.NewApi(app, vehicleService)

	app.Listen(":3000")
}
