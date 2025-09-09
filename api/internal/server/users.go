package server

import (
	"context"
	"net/http"

	"callisto/internal/request"
)

func (app *Application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.db.GetAllUsers(context.Background())
	if err != nil {
		request.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	request.JSON(w, http.StatusOK, map[string]any{"users": users})
}
