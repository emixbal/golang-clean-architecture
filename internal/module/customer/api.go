package customer

import (
	"context"
	"golang-clean-architecture/domain"
	"golang-clean-architecture/internal/util"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type api struct {
	customerService domain.CustomerService
}

func NewApi(app *fiber.App, customerService domain.CustomerService) {
	api := api{
		customerService: customerService,
	}

	app.Get("/", api.FooBar)
	app.Get("/v1/customers", api.AllCustomers)
}

func (a api) FooBar(c *fiber.Ctx) error {
	log.Println("ok")
	return c.Status(200).JSON(fiber.Map{"message": "ok"})
}

func (a api) AllCustomers(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	apiResponse := a.customerService.All(c)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.JSON(apiResponse)
}
