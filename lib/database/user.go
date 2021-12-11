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
	GetAccountByEmailNullToken(login model.User_Login) (user model.User, err error)
	GetAccountByEmailWithToken(login model.User_Login) (user model.User, err error)
	UpdateToken(token string, id int) error
}

func NewUserModel(db *sql.DB) *UserModelImplementation {
	return &UserModelImplementation{db: db}
}

func (um *UserModelImplementation) AddUser(newUser model.User) (user model.User, err error) {
	tx, err := um.db.Begin()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return model.User{}, err
	}
	res, err := tx.Exec(
		"insert into users(name,email,password,created_at,updated_at) values (?,?,?,?,?)",
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
		fmt.Println(err)

		return model.User{}, errors.New("insert failed")
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return user, err
	}
	row := tx.QueryRow("select * from users where id = ?", insertID)
	var tokenNullString sql.NullString
	var createdAtString string
	var updatedAtString string
	var deletedAtString sql.NullTime
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &tokenNullString,
		&createdAtString, &updatedAtString, &deletedAtString)
	if err != nil {
		tx.Rollback()
		fmt.Println(err, 123)

		return model.User{}, err
	}
	user.Created_at, _ = helper.DecodeTime(createdAtString)
	user.Updated_at, _ = helper.DecodeTime(updatedAtString)
	tx.Commit()
	return user, nil
}

func (um *UserModelImplementation) GetAccountByEmailNullToken(login model.User_Login) (user model.User, err error) {
	tx, err := um.db.Begin()
	if err != nil {
		tx.Rollback()
		return user, err
	}
	row := tx.QueryRow("select * from users where email = ? and deleted_at is null", login.Email)
	var createdAtString string
	var updatedAtString string
	var deletedAtString sql.NullString
	var nullToken sql.NullString
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &nullToken,
		&createdAtString, &updatedAtString, &deletedAtString)
	if err != nil {
		tx.Rollback()
		return model.User{}, err
	}
	user.Created_at, _ = helper.DecodeTime(createdAtString)
	user.Updated_at, _ = helper.DecodeTime(updatedAtString)
	tx.Commit()
	return user, nil
}

func (um *UserModelImplementation) GetAccountByEmailWithToken(login model.User_Login) (user model.User, err error) {
	tx, err := um.db.Begin()
	if err != nil {
		tx.Rollback()
		return user, err
	}
	row := tx.QueryRow("select * from users where email = ? and deleted_at is null", login.Email)
	var createdAtString string
	var updatedAtString string
	var deletedAtString sql.NullString
	// var nullToken sql.NullString
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Token,
		&createdAtString, &updatedAtString, &deletedAtString)
	if err != nil {
		tx.Rollback()
		return model.User{}, err
	}
	user.Created_at, _ = helper.DecodeTime(createdAtString)
	user.Updated_at, _ = helper.DecodeTime(updatedAtString)
	tx.Commit()
	return user, nil
}

func (um *UserModelImplementation) UpdateToken(token string, id int) error {
	tx, err := um.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec("update users set token = ? where id = ?", token, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affected < 1 {
		tx.Rollback()
		return errors.New("update token failed")
	}
	tx.Commit()
	return nil
}
