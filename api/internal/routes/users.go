package routes

import (
	"net/http"

	"jinji/internal/database"
	"jinji/internal/response"
)

func getUsers(w http.ResponseWriter, r *http.Request) {

	// users, _ := database.GetAllUsers()

	response.JSON(w, http.StatusOK, map[string]any{
		"status":  "success",
		"message": "fetched all users",
		"users":   []database.User{},
	})
}
