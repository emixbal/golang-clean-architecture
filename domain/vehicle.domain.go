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
	FindHistorical(ctx context.Context, vin string) ApiResponse
	StoreHistorical(ctx context.Context, request VehicleHistoricalRequest) ApiResponse
}

type VehicleHistorical struct {
	ID        int64         `json:"id"`
	VIN       string        `json:"vin"`
	Brand     string        `json:"brand"`
	Histories []HistoryData `json:"histories"`
}

type VehicleHistoricalRequest struct {
	VIN      string `json:"vin"`
	Brand    string `json:"brand"`
	PIC      string `json:"pic"`
	Notes    string `json:"notes"`
	Customer int64  `json:"customer"`
}
