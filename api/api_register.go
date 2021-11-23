package api

import (
	"net/http"
)

// RegisterJSON congtains JSON param for user registration
type RegisterJSON struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RegisterHandler handle user registration API endpoint
type RegisterHandler Handler

func (h RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var jsonParams RegisterJSON

	if err := ReadJSON(r, &jsonParams); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.AddNewUser(jsonParams.Name, jsonParams.Email, jsonParams.Password); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, DefaultResponse{
		Success: true,
		Message: "user created successfully",
	})
}
