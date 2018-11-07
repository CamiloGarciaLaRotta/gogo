package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"gogo/config"

	_ "github.com/lib/pq"
)

// GadgetService is the postgres implementation of a gogo.GadgetService
type GadgetService struct {
	DB *sql.DB
}

// New returns a postgres Gadget service
func New(cfg config.Postgres) (*GadgetService, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, errors.Wrap(err, "could not open DB connection")
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, errors.Wrap(err, "could not ping DB")
	}

	return &GadgetService{db}, nil
}
