package dbm_test

import (
	"hlaing-backend-test-1/dbm"
	"hlaing-backend-test-1/env"
	"math/rand"
	"os"
	"testing"
	"time"
)

var (
	dbManager *dbm.DatabaseManager
	envars    = env.GetEnvironment()
)

func TestMain(m *testing.M) {
	envars.TablePrefix = "test_dbm_"
	dbManager = dbm.NewDatabaseManager(envars)

	dbManager.TearDown()
	dbManager.Setup()

	rand.Seed(time.Now().UTC().UnixNano())
	status := m.Run()

	if status == 0 {
		dbManager.TearDown()
	}
	os.Exit(status)
}
