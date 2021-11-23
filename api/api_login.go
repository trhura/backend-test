package api

import (
	"net/http"

	"gitea.com/go-chi/session"
)

// LoginHandler handle user login API endpoint
type LoginHandler Handler

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var username, password, ok = r.BasicAuth()
	if !ok {
		WriteError(w, http.StatusBadRequest, "unable to parse basic auth")
		return
	}

	if _, err := h.AuthenticateUser(username, password); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var store = session.GetSession(r)
	if err := store.Set(`authenticated_user`, username); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, DefaultResponse{
		Success: true,
		Message: "user login successful",
	})
}
