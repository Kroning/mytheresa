package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Kroning/mytheresa/internal/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type NodeConfig struct {
	Host           string
	Port           string
	User           string
	Password       string
	Database       string
	MaxOpen        uint
	Timeout        time.Duration
	MigrationsPath string
}

type Config struct {
	Master NodeConfig
}

type Storage struct {
	Master *sqlx.DB
}

func New(cfg Config) (*Storage, error) {
	db, err := sql.Open("pgx", connectionString(cfg.Master))
	if err != nil {
		return nil, fmt.Errorf("master DB: %v", err)
	}
	masterDB := sqlx.NewDb(db, "pgx")
	masterDB.DB.SetMaxOpenConns(int(cfg.Master.MaxOpen))

	s := &Storage{
		Master: masterDB,
	}

	err = s.MigrateUp(cfg.Master.MigrationsPath)
	if err != nil {
		logger.Error(context.Background(), "migrations error", zap.Error(err))
		return nil, fmt.Errorf("migration error: %v", err)
	}

	return s, err
}

func connectionString(c NodeConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Database)
}

func (s *Storage) Close() error {
	if err := s.Master.Close(); err != nil {
		return errors.Wrap(err, "db master: %v")
	}

	return nil
}

func (s *Storage) MigrateUp(migrationsPath string) error {
	var err error
	driver, err := postgres.WithInstance(s.Master.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	var m *migrate.Migrate

	m, err = migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}
	return nil
}
