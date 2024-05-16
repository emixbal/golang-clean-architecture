package vehicle

import (
	"context"
	"golang-clean-architecture/domain"
	"golang-clean-architecture/internal/util"
	"time"

	"github.com/gofiber/fiber/v2"
)

type api struct {
	vehicleService domain.VehicleService
}

func NewApi(app *fiber.App, vehicleService domain.VehicleService) {
	api := api{
		vehicleService: vehicleService,
	}

	app.Get("/v1/vehicles/histories", api.getVehicleHistories)
	app.Post("/v1/vehicles/histories", api.addVehicleHistory)
}

func (a api) getVehicleHistories(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	vin := ctx.Query("vin")
	if vin == "" {
		apiResponse := domain.ApiResponse{
			Code:    "400",
			Message: "invalid param. vin is required",
		}
		util.ResponseInterceptor(c, &apiResponse)
		return ctx.Status(400).JSON(apiResponse)
	}

	apiResponse := a.vehicleService.FindHistorical(c, vin)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.JSON(apiResponse)
}

func (a api) addVehicleHistory(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	var request domain.VehicleHistoricalRequest
	if err := ctx.BodyParser(&request); err != nil {
		apiResponse := domain.ApiResponse{
			Code:    "400",
			Message: "invalid body",
		}
		util.ResponseInterceptor(c, &apiResponse)
		return ctx.Status(400).JSON(apiResponse)
	}

	apiResponse := a.vehicleService.StoreHistorical(c, request)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(200).JSON(apiResponse)
}
