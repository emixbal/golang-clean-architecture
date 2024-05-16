package domain

import (
	"context"
	"time"
)

type HistoryDetail struct {
	ID         int64     `db:"id"`
	VehicleID  int64     `db:"history_id"`
	CustomerID int64     `db:"customer_id"`
	Notes      string    `db:"notes"`
	PIC        string    `db:"pic"`
	CreatedAt  time.Time `db:"created_at"`
}

type HistoryDetailRepository interface {
	FindByID(ctx context.Context, id int64) (HistoryDetail, error)
	FindByVehicleID(ctx context.Context, vehicle_id int64) ([]HistoryDetail, error)
	Insert(ctx context.Context, historyDetail *HistoryDetail) error
}

type HistoryDetailService interface{}
