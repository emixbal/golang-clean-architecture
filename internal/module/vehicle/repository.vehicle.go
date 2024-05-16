package vehicle

import (
	"context"
	"database/sql"
	"golang-clean-architecture/domain"

	"github.com/doug-martin/goqu/v9"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.VehicleRepository {
	return &repository{
		db: goqu.New("default", con),
	}
}

// FindById implements domain.VehicleRepository.
func (r *repository) FindById(ctx context.Context, id string) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicles").Where(goqu.Ex{
		"id": id,
	})

	if _, err := dataset.ScanStructContext(ctx, &vehicle); err != nil {
		return domain.Vehicle{}, err
	}

	return
}

// FindByVIN implements domain.VehicleRepository.
func (r *repository) FindByVIN(ctx context.Context, vin string) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicles").Where(goqu.Ex{
		"vin": vin,
	})

	if _, err := dataset.ScanStructContext(ctx, &vehicle); err != nil {
		return domain.Vehicle{}, err
	}

	return
}

// Insert implements domain.VehicleRepository.
func (r *repository) Insert(ctx context.Context, vehicle *domain.Vehicle) error {
	executor := r.db.Insert("vehicles").
		Rows(*vehicle).
		Returning("id").
		Executor()

	_, err := executor.ScanStructContext(ctx, &vehicle)

	return err
}
