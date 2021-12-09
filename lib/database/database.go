package database

import (
	"database/sql"
	"errors"
	"fmt"

	"test-project-1/helper"
	"test-project-1/model"
)

type UserModelImplementation struct {
	db *sql.DB
}

type UserModel interface {
	AddUser(newUser model.User) (user model.User, err error)
}

func NewUserModel(db *sql.DB) *UserModelImplementation {
	return &UserModelImplementation{db: db}
}

func (um *UserModelImplementation) AddUser(newUser model.User) (user model.User, err error) {
	tx, err := um.db.Begin()
	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}
	res, err := tx.Exec("insert into users(name,email,password,created_at,updated_at) values (?,?,?,?,?)",
		newUser.Name, newUser.Email, newUser.Password,
		newUser.Created_at, newUser.Updated_at,
	)

	if err != nil {
		tx.Rollback()
		fmt.Println(err)

		return model.User{}, err
	}

	affected, err := res.RowsAffected()

	if err != nil {
		tx.Rollback()
		fmt.Println(err)

		return model.User{}, err
	}

	if affected < 1 {
		tx.Rollback()
		return model.User{}, errors.New("insert failed")
	}

	row := tx.QueryRow("select * from users")
	var tokenNullString sql.NullString
	var createdAtString string
	var updatedAtString string
	var deletedAtString sql.NullString
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &tokenNullString,
		&createdAtString, &updatedAtString, &deletedAtString)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return model.User{}, err
	}
	user.Created_at, _ = helper.DecodeTime(createdAtString)
	user.Updated_at, _ = helper.DecodeTime(updatedAtString)
	tx.Commit()
	return user, nil
}
