package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/otaviopontes/api-go/src/controllers"
	"github.com/otaviopontes/api-go/src/models"
	"github.com/otaviopontes/api-go/src/repositories"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockController := controllers.NewMockUserController(ctrl)
	mockRepo := repositories.NewMockUserRepository(ctrl)

	mockUser := models.User{
		Id:       uuid.New(),
		Email:    "user@example.com",
		Password: "hashedpassword",
	}

	mockRepo.EXPECT().
		SearchByEmail("user@example.com").
		Return(mockUser, nil).
		Times(1)

	body := bytes.NewBuffer([]byte(`{"email": "user@example.com", "password": "password123"}`))
	req, _ := http.NewRequest("POST", "/login", body)
	w := httptest.NewRecorder()

	mockController.Login(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseBody map[string]interface{}
	json.NewDecoder(w.Body).Decode(&responseBody)
	assert.NotEmpty(t, responseBody["token"])
	assert.Equal(t, "mockedToken", responseBody["token"])
}
