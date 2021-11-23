package api

import (
	"net/http"

	"gitea.com/go-chi/session"
)

// LogoutHandler handle user registration API endpoint
type LogoutHandler Handler

func (h LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var store = session.GetSession(r)

	if err := store.Flush(); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := store.Destroy(w, r); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, DefaultResponse{
		Success: true,
		Message: "user logout successful",
	})
}
