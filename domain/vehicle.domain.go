package domain

import (
	"context"
	"time"
)

type Vehicle struct {
	ID        int64     `db:"id" goqu:"omitempty"`
	VIN       string    `db:"vin"`
	Brand     string    `db:"brand"`
	CreatedAt time.Time `db:"created_at"`
}

type VehicleRepository interface {
	FindById(ctx context.Context, id string) (Vehicle, error)
	FindByVIN(ctx context.Context, vin string) (Vehicle, error)
	Insert(ctx context.Context, vehicle *Vehicle) error
	// FindDetailHistories(ctx context.Context, id int64) ([]HistoryDetail, error)
}

type VehicleService interface {
	FindByID(ctx context.Context, vehicle Vehicle) ApiResponse
}
