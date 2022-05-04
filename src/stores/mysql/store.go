package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"dkfbasel.ch/covid-evidence/logger"
	"dkfbasel.ch/covid-evidence/stores"
)

// Store is the struct to fullfil the store interface
type Store struct {
	DB     *sqlx.DB
	DryRun bool
}

// Setup the database
func (s *Store) Setup(databaseURL string, dryRun bool) {
	if dryRun {
		s.DryRun = true
	}
	db, err := sqlx.Connect("mysql", databaseURL)
	if err != nil {
		logger.Fatal("could not connect to mysql db", logger.Err(err))
	}
	logger.Info("database connected")
	s.DB = db
}

// Close ...
func (s *Store) Close() {
	s.DB.Close()
}

// FetchBasic ...
func (s *Store) FetchBasic(sources ...string) ([]stores.Record, stores.Index, error) {
	return s.fetchCoveBasic(sources...)
}
