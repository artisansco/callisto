package routes

import (
	"net/http"

	"jinji/internal/request"
	"jinji/internal/response"
	"jinji/internal/validator"
)

func generateToken(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	request.DecodeJSON(w, r, payload)

	if validator.IsEmail(payload.Email) || len(payload.Password) == 0 {
		response.JSON(w, http.StatusBadRequest, map[string]string{
			"status":  "fail",
			"message": "Email and password are required",
		})
		return
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"email": payload.Email,
	// 	"exp":   time.Now().Add(time.Hour * 24).Unix(),
	// })
	// tokenString, err := token.SignedString([]byte("secret"))
	// if err != nil {
	// 	request.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	// 	return
	// }

	response.JSON(w, http.StatusOK, map[string]any{
		"status":  "success",
		"message": "JWT created successfully",
		"data": map[string]string{
			"token": "01992ed4-98ae-7a81-be8e-e4e9fc5b922a",
		},
	})
}
