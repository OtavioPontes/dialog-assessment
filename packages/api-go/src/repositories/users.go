package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/otaviopontes/api-go/src/models"
)

type UserRepository interface {
	Create(user models.User) error
	Get(nameOrNick string) ([]models.User, error)
	GetById(userId uuid.UUID) (models.User, error)
	Update(userId uuid.UUID, user models.User) error
	Delete(userId uuid.UUID) error
	SearchByEmail(email string) (models.User, error)
	SearchPassword(id uuid.UUID) (string, error)
	UpdatePassword(userId uuid.UUID, password []byte) error
}

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository *Users) Create(user models.User) error {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nick, email, password) VALUES ($1, $2, $3, $4);",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (repository *Users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (repository *Users) GetById(userId uuid.UUID) (models.User, error) {

	lines, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = $1",
		userId,
	)
	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}

	} else {
		return models.User{}, errors.New("user not found with this id")
	}

	return user, nil

}

func (repository *Users) Update(userId uuid.UUID, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = $1, nick = $2, email = $3 where id = $4",
	)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nick, user.Email, userId)
	if err != nil {
		return err
	}

	return nil
}

func (repository *Users) Delete(userId uuid.UUID) error {
	statement, err := repository.db.Prepare(
		"delete from users where id = $1",
	)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(userId)
	if err != nil {
		return err
	}

	return nil
}

func (repository *Users) SearchByEmail(email string) (models.User, error) {
	line, err := repository.db.Query("select id, password from users where email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		err := line.Scan(
			&user.Id,
			&user.Password,
		)
		if err != nil {
			return models.User{}, err
		}
	} else {
		return models.User{}, errors.New("user not found with this email")
	}

	return user, nil
}

func (repository *Users) SearchPassword(id uuid.UUID) (string, error) {
	line, err := repository.db.Query("select password from users where id = $1", id)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var password string
	if line.Next() {
		err := line.Scan(
			&password,
		)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("user not found with this id")
	}

	return password, nil
}

func (repository *Users) UpdatePassword(userId uuid.UUID, password []byte) error {
	statement, err := repository.db.Prepare(
		"update users set password = $1 where id = $2",
	)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(password, userId)
	if err != nil {
		return err
	}

	return nil
}
