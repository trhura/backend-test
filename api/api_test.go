package api_test

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"

	"io"
	"os"
	"testing"
	"time"

	"hlaing-backend-test-1/api"
	"hlaing-backend-test-1/dbm"
	"hlaing-backend-test-1/env"

	"github.com/go-chi/chi/v5"
)

var (
	envars    = env.GetEnvironment()
	dbManager *dbm.DatabaseManager
	router    *chi.Mux
)

func TestMain(m *testing.M) {
	envars.TablePrefix = "test_api_"
	dbManager = dbm.NewDatabaseManager(envars)
	router = api.NewAPIRouter(dbManager)

	dbManager.TearDown()
	dbManager.Setup()

	rand.Seed(time.Now().UTC().UnixNano())
	status := m.Run()

	if status == 0 {
		dbManager.TearDown()
	}
	os.Exit(status)
}

// CallTestAPI is Helper function to test a handler with params json body and
// save resulting JSON output in `outjson` interface
func CallTestAPI(method, path string, injson, outjson interface{}, headers map[string]string) *http.Response {
	var err error
	var body io.Reader
	var data []byte

	if injson != nil {
		if data, err = json.Marshal(injson); err != nil {
			panic(err)
		}
		body = bytes.NewReader(data)
	}

	var recorder = httptest.NewRecorder()
	var request = httptest.NewRequest(method, path, body)

	for header, value := range headers {
		request.Header.Set(header, value)
	}

	router.ServeHTTP(recorder, request)

	if outjson != nil {
		data = recorder.Body.Bytes()
		if err = json.Unmarshal(data, outjson); err != nil {
			panic(err)
		}
	}

	return recorder.Result()
}
