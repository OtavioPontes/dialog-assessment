package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/otaviopontes/api-go/src/authentication"
	"github.com/otaviopontes/api-go/src/database"
	"github.com/otaviopontes/api-go/src/models"
	"github.com/otaviopontes/api-go/src/repositories"
	"github.com/otaviopontes/api-go/src/responses"
	"github.com/otaviopontes/api-go/src/security"
)

// @Summary      Create a new user
// @Description  Allows for the creation of a new user in the system.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User data"
// @Success      201
// @Failure      422  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare(true); err != nil {
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
	err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)

}

// @Summary      Get a user by ID
// @Description  Retrieves a user from the system by their ID.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
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

	user, err := repository.GetById(userId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// @Summary      Update user details
// @Description  Updates the information of a user in the system.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      string       true  "User ID"
// @Param        user  body      models.User  true  "Updated user data"
// @Success      204
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      401  {object}  responses.ErrorResponse
// @Failure      403  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdFromToken, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdFromToken {
		responses.Error(w, http.StatusForbidden, errors.New("it is not possible to update a user if not yours"))
		return
	}

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

	if err = user.Prepare(false); err != nil {
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

	err = repository.Update(userId, user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// @Summary      Delete a user
// @Description  Deletes a user from the system by their ID.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      204
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      401  {object}  responses.ErrorResponse
// @Failure      403  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdFromToken, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdFromToken {
		responses.Error(w, http.StatusForbidden, errors.New("it is not possible to delete a user if not yours"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	err = repository.Delete(userId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UpdatePassword godoc
// @Summary      Update user password
// @Description  Allows a user to update their password in the system.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id        path      string                 true  "User ID"
// @Param        password  body      map[string]string      true  "Password data"
// @Success      204
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      401  {object}  responses.ErrorResponse
// @Failure      403  {object}  responses.ErrorResponse
// @Failure      422  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /users/{id}/password [put]
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdFromToken, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdFromToken {
		responses.Error(w, http.StatusForbidden, errors.New("it is not possible to change other user's password"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	password := struct {
		New     string `json:"new"`
		Current string `json:"current"`
	}{}

	if err := json.Unmarshal(body, &password); err != nil {
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

	savedPassword, err := repository.SearchPassword(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(password.Current, savedPassword); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("it is not possible to change other user's password"))
		return
	}

	hashedPassword, err := security.Hash(password.New)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	err = repository.UpdatePassword(userId, hashedPassword)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
