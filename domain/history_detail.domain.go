package domain

import "time"

type HistoryDetail struct {
	ID        int64     `db:"id"`
	HistoryID int64     `db:"history_id"`
	Notes     string    `db:"notes"`
	PIC       string    `db:"pic"`
	CreatedAt time.Time `db:"created_at"`
}

type HistoryDetailRepository interface{}

type HistoryDetailService interface{}
