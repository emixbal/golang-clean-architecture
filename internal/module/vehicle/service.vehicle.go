package vehicle

import (
	"context"
	"golang-clean-architecture/domain"
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
