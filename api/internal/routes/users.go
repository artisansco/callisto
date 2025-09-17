package routes

import (
	"context"
	"net/http"

	"jinji/internal/response"
)

func (app *Application) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.db.GetAllUsers(context.Background())
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to fetch users",
		})
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"status":  "success",
		"message": "fetched all users",
		"users":   users,
	})
}
