package repositories_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/google/uuid"
	"github.com/otaviopontes/api-go/src/models"
	"github.com/otaviopontes/api-go/src/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)

	user := models.User{
		Name:     "John Doe",
		Nick:     "johnd",
		Email:    "john@example.com",
		Password: "password123",
	}

	mock.ExpectPrepare("INSERT INTO users").
		ExpectExec().
		WithArgs(user.Name, user.Nick, user.Email, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepo.Create(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userId := uuid.New()

	rows := sqlmock.NewRows([]string{"id", "name", "nick", "email", "createdAt"}).
		AddRow(userId, "John Doe", "johnd", "john@example.com", time.Now())

	mock.ExpectQuery("select id, name, nick, email, createdAt from users where id =").
		WithArgs(userId).
		WillReturnRows(rows)

	user, err := userRepo.GetById(userId)
	assert.NoError(t, err)
	assert.Equal(t, userId, user.Id)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "johnd", user.Nick)
	assert.Equal(t, "john@example.com", user.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSearchByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)

	email := "john@example.com"
	userId := uuid.New()
	password := "hashedpassword"

	rows := sqlmock.NewRows([]string{"id", "password"}).
		AddRow(userId, password)

	mock.ExpectQuery("select id, password from users where email =").
		WithArgs(email).
		WillReturnRows(rows)

	user, err := userRepo.SearchByEmail(email)
	assert.NoError(t, err)
	assert.Equal(t, userId, user.Id)
	assert.Equal(t, password, user.Password)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)

	userId := uuid.New()
	user := models.User{
		Name:  "John Doe",
		Nick:  "johnd",
		Email: "john@example.com",
	}

	mock.ExpectPrepare("update users set name =").
		ExpectExec().
		WithArgs(user.Name, user.Nick, user.Email, userId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepo.Update(userId, user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)

	userId := uuid.New()

	mock.ExpectPrepare("delete from users where id =").
		ExpectExec().
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepo.Delete(userId)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)

	userId := uuid.New()
	password := []byte("newhashedpassword")

	mock.ExpectPrepare("update users set password =").
		ExpectExec().
		WithArgs(password, userId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepo.UpdatePassword(userId, password)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
