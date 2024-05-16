package vehicle

import (
	"context"
	"golang-clean-architecture/domain"
	"log"
	"time"
)

type service struct {
	vehicleRepository domain.VehicleRepository
	historyRepository domain.HistoryRepository
}

func NewService(
	vehicleRepository domain.VehicleRepository,
	historyRepository domain.HistoryRepository,
) domain.VehicleService {
	return &service{
		vehicleRepository: vehicleRepository,
		historyRepository: historyRepository,
	}
}

// FindHistorical implements domain.VehicleService.
func (s *service) FindHistorical(ctx context.Context, vin string) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, vin)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	if vehicle == (domain.Vehicle{}) {
		return domain.ApiResponse{
			Code:    "200",
			Message: "vehicle not found",
		}
	}

	histories, err := s.historyRepository.FindByVehicleID(ctx, vehicle.ID)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	var historiesData []domain.HistoryData
	for _, val := range histories {
		historiesData = append(historiesData, domain.HistoryData{
			VehicleID:  val.VehicleID,
			CustomerID: val.CustomerID,
			Notes:      val.Notes,
			PIC:        val.PIC,
			Date:       val.CreatedAt.Format(time.RFC822Z),
		})
	}

	result := domain.VehicleHistorical{
		ID:        vehicle.ID,
		VIN:       vehicle.VIN,
		Brand:     vehicle.Brand,
		Histories: historiesData,
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "ok",
		Data:    result,
	}
}

// StoreHistorical implements domain.VehicleService.
func (s *service) StoreHistorical(ctx context.Context, request domain.VehicleHistoricalRequest) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, request.VIN)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: "err.Error()",
		}
	}

	if vehicle == (domain.Vehicle{}) {
		vehicle.Brand = request.Brand
		vehicle.VIN = request.VIN
		vehicle.CreatedAt = time.Now()

		err = s.vehicleRepository.Insert(ctx, &vehicle)
		if err != nil {
			return domain.ApiResponse{Code: "500", Message: err.Error()}
		}
	}

	err = s.historyRepository.Insert(ctx, &domain.History{
		VehicleID:  vehicle.ID,
		CustomerID: request.Customer,
		PIC:        request.PIC,
		Notes:      request.Notes,
	})

	if err != nil {
		log.Println(err)
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "ok",
	}
}
