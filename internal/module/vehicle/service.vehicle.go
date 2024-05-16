package vehicle

import (
	"context"
	"golang-clean-architecture/domain"
)

type service struct {
	vehicleRepository domain.VehicleRepository
}

func NewService(vehicleRepository domain.VehicleRepository) domain.VehicleService {
	return &service{vehicleRepository: vehicleRepository}
}

// FindByID implements domain.VehicleService.
func (s *service) FindByID(ctx context.Context, vehicle domain.Vehicle) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindById(ctx, "a")
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "success",
		Data:    vehicle,
	}
}
