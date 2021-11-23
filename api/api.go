package api

import (
	"fmt"
	"hlaing-backend-test-1/dbm"
	"log"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"gitea.com/go-chi/session"
	_ "gitea.com/go-chi/session/redis"

	"github.com/go-chi/chi/v5"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	validate   = validator.New()
	cookieName = "ssn"
)

// Handler is a generic http handler with wraps database operations
type Handler struct {
	http.Handler
	*dbm.DatabaseManager
}

// DefaultResponse contains default JSON response for API
type DefaultResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// NewAPIRouter return a new api router
func NewAPIRouter(dbManager *dbm.DatabaseManager) *chi.Mux {
	r := chi.NewRouter()

	r.Use(session.Sessioner(session.Options{
		Provider:       "redis",
		ProviderConfig: "addr:" + dbManager.RedisAddr + ",prefix=session:",
		CookieName:     cookieName,
		CookieLifeTime: int(dbManager.SessionLifetime.Seconds()),
	}))

	r.Method("POST", "/register", RegisterHandler{DatabaseManager: dbManager})
	r.Method("POST", "/login", LoginHandler{DatabaseManager: dbManager})
	r.Method("POST", "/logout", LogoutHandler{DatabaseManager: dbManager})

	return r
}

// ReadJSON read JSON from http request into `out`
func ReadJSON(r *http.Request, out interface{}) error {
	var err error
	var bytes []byte

	if bytes, err = ioutil.ReadAll(r.Body); err != nil {
		return err
	}

	if len(bytes) <= 0 {
		return fmt.Errorf("empty request body")
	}

	if err = json.Unmarshal(bytes, &out); err != nil {
		return err
	}

	if err = validate.Struct(out); err != nil {
		return err
	}

	return nil
}

// WriteError write json response according to err
func WriteError(w http.ResponseWriter, status int, messsage string) {
	w.WriteHeader(status)
	WriteJSON(w, DefaultResponse{
		Success: false,
		Message: messsage,
	})
}

// WriteJSON generate JSON response in { "sucess": ..., } format
func WriteJSON(w http.ResponseWriter, response interface{}) {
	var err error
	var bytes []byte
	if bytes, err = json.Marshal(response); err != nil {
		log.Fatalln("WriteJSON: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
