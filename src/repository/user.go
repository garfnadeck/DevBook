package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

// Open database do the User table
func UserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create user
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name,nick,email,password) VALUES (?,?,?,?)",
	)
	if err != nil {
		return 0, nil
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	lastIdInsert, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastIdInsert), nil
}

// Search for all users
func (repository Users) SearchUser(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	queryNameOrNick, err := repository.db.Query("select id, name, nick, email, createdAt from users where name LIKE ? or nick Like ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer queryNameOrNick.Close()

	var users []models.User

	for queryNameOrNick.Next() {
		var user models.User
		if err = queryNameOrNick.Scan(
			&user.ID,
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

func (repository Users) SearchById(ID uint64) (models.User, error) {
	query, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer query.Close()

	var user models.User

	if query.Next() {
		if err = query.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
