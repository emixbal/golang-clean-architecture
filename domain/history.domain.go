package domain

import (
	"context"
	"time"
)

type History struct {
	ID         int64     `db:"id" goqu:"omitempty"`
	VehicleID  int64     `db:"vehicle_id"`
	CustomerID int64     `db:"customer_id"`
	Notes      string    `db:"notes"`
	PIC        string    `db:"pic"`
	CreatedAt  time.Time `db:"created_at"`
}

type HistoryRepository interface {
	FindByID(ctx context.Context, id int64) (History, error)
	FindByVehicleID(ctx context.Context, vehicle_id int64) ([]History, error)
	Insert(ctx context.Context, history *History) error
}

type HistoryService interface{}

type HistoryData struct {
	VehicleID  int64  `json:"vehicle_id"`
	CustomerID int64  `json:"customer_id"`
	Notes      string `json:"notes"`
	PIC        string `json:"pic"`
	Date       string `json:"date"`
}
