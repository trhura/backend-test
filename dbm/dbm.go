package dbm

import (
	"hlaing-backend-test-1/env"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DatabaseManager wraps all database operations
type DatabaseManager struct {
	*sqlx.DB
	*UserManager
}

// NewDatabaseManager construct a new database manager
func NewDatabaseManager(envars *env.Envars) *DatabaseManager {
	db, err := sqlx.Connect("mysql", envars.DatabaseURL)
	if err != nil {
		log.Fatalln(err)
	}

	return &DatabaseManager{
		DB:          db,
		UserManager: NewUserManager(envars, db),
	}
}

// Setup is used for testing setup
func (dm *DatabaseManager) Setup() {
	dm.UserManager.Setup()
}

// TearDown is used for testing teardown
func (dm *DatabaseManager) TearDown() {
	dm.UserManager.TearDown()
}
