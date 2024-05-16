package history

import (
	"context"
	"database/sql"
	"golang-clean-architecture/domain"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.HistoryRepository {
	return &repository{
		db: goqu.New("default", con),
	}
}

// FindByID implements domain.HistoryRepository.
func (*repository) FindByID(ctx context.Context, id int64) (domain.History, error) {
	panic("unimplemented")
}

// FindByVehicleID implements domain.HistoryRepository.
func (r *repository) FindByVehicleID(ctx context.Context, vehicle_id int64) (histories []domain.History, err error) {
	dataset := r.db.From("histories").
		Where(goqu.Ex{
			"vehicle_id": vehicle_id,
		}).
		Order(goqu.I("id").Asc())

	// Pass a pointer to the slice
	err = dataset.ScanStructsContext(ctx, &histories)

	return
}

// Insert implements domain.HistoryRepository.
func (r *repository) Insert(ctx context.Context, history *domain.History) error {
	executor := r.db.Insert("histories").
		Rows(goqu.Record{
			"vehicle_id":  history.VehicleID,
			"customer_id": history.CustomerID,
			"notes":       history.Notes,
			"pic":         history.PIC,
			"created_at":  time.Now(),
		}).
		Returning("id").Executor()

	_, err := executor.Exec()
	return err
}
