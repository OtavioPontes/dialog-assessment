package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/otaviopontes/api-go/src/authentication"
	"github.com/otaviopontes/api-go/src/database"
	"github.com/otaviopontes/api-go/src/models"
	"github.com/otaviopontes/api-go/src/repositories"
	"github.com/otaviopontes/api-go/src/responses"
	"github.com/otaviopontes/api-go/src/security"
	_ "github.com/swaggo/http-swagger"
)

// @Summary      User Login
// @Description  Authenticate a user with their email and password.
// @Tags         Login
// @Accept       json
// @Produce      json
// @Param        email     body  string  true  "User email"
// @Param        password  body  string  true  "User password"
// @Success      200   {object}  responses.AuthResponse
// @Failure      400   {object}  responses.ErrorResponse
// @Failure      401   {object}  responses.ErrorResponse
// @Failure      404   {object}  responses.ErrorResponse
// @Failure      422   {object}  responses.ErrorResponse
// @Failure      500   {object}  responses.ErrorResponse
// @Router       /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	savedUser, err := repository.SearchByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusNotFound, err)
		return
	}

	if err := security.VerifyPassword(user.Password, savedUser.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(savedUser.Id)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userId := savedUser.Id.String()

	responses.JSON(w, http.StatusOK, responses.AuthResponse{
		Id:    userId,
		Token: token,
	})

}
