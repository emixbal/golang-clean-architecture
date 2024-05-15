package domain

import (
	"context"
	"time"
)

type History struct {
	ID        int64     `db:"id"`
	NoRangka  string    `db:"no_rengka"`
	Merek     string    `db:"merek"`
	CreatedAt time.Time `db:"created_at"`
}

type HistoryRepository interface {
	FindById(ctx context.Context, id int64) (History, error)
	FindByNoRangka(ctx context.Context, id int64) (History, error)
	FindDetailHistories(ctx context.Context, id int64) ([]HistoryDetail, error)
	Insert(ctx context.Context, history *History) error
	InsertDetail(ctx context.Context, detail *HistoryDetail) error
}

type HistoryService interface{}
