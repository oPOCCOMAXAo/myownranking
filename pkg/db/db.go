package db

import (
	"log/slog"

	slogGorm "github.com/orandin/slog-gorm"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DSN string `env:"DSN,required"`
}

func NewPostgres(
	cfg Config,
	logger *slog.Logger,
) (*gorm.DB, error) {
	gormCfg := &gorm.Config{}

	if logger != nil {
		gormCfg.Logger = slogGorm.New(
			slogGorm.WithHandler(logger.Handler()),
		)
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN), gormCfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}
