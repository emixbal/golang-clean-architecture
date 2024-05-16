package domain

import (
	"context"
	"time"
)

type Vehicle struct {
	ID        int64     `db:"id"`
	VIN       string    `db:"vin"`
	Brand     string    `db:"brand"`
	CreatedAt time.Time `db:"created_at"`
}

type VehicleRepository interface {
	FindById(ctx context.Context, id int64) (Vehicle, error)
	FindByNoRangka(ctx context.Context, id int64) (Vehicle, error)
	FindDetailHistories(ctx context.Context, id int64) ([]HistoryDetail, error)
	Insert(ctx context.Context, history *Vehicle) error
	InsertDetail(ctx context.Context, detail *HistoryDetail) error
}

type VehicleService interface{}
