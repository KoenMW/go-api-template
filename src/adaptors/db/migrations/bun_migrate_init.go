package migration

import (
	"context"
	"go-api/domain/model"
	"go-api/ports/db"
	"go-api/ports/migration"
)

type Migrations struct {
	dbAdapter db.DB
}

func (m *Migrations) Up(ctx context.Context) error {
	db, err := m.dbAdapter.GetDB(ctx)
	if err != nil {
		return err
	}

	if _, err := db.NewCreateTable().
		Model((*model.User)(nil)).
		IfNotExists().
		ForeignKey(`("role_id") REFERENCES roles(id)`).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func NewMigrations(dbAdapter db.DB) migration.Migration {
	migrations := &Migrations{
		dbAdapter: dbAdapter,
	}

	dbAdapter.NotifyFirstAvailable(migrations.Up)
	return migrations
}
