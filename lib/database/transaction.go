package database

import (
	"database/sql"
	"fmt"
	"log"
	"test-project-1/model"
)

type TransactionModelImplementation struct {
	db *sql.DB
}

type TransactionModel interface {
	addTransactionDetail(tx *sql.Tx, t model.Transaction_Detail) (err error)
	CreateTransaction(t model.Transaction, d []model.Transaction_Detail) (id int, err error)
	InsertTransactionID(id int) (err error)
	GetTransactionDetailByTransactionID(id int) (details []model.Transaction_Detail, err error)
}

func NewTransactionModel(db *sql.DB) *TransactionModelImplementation {
	return &TransactionModelImplementation{db: db}
}

func (tm *TransactionModelImplementation) addTransactionDetail(tx *sql.Tx, t model.Transaction_Detail) (err error) {
	res, err := tx.Exec("insert into transaction_details(item_id, item_quantity, item_price, item_cost, created_at, updated_at) values(?,?,?,?,?,?)",
		t.ItemID, t.ItemQuantity, t.ItemPrice, t.ItemCost, t.Created_at, t.Updated_at)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	if affected < 1 {
		tx.Rollback()
		log.Println(err)
		return err
	}
	return nil
}

func (tm *TransactionModelImplementation) CreateTransaction(t model.Transaction, d []model.Transaction_Detail) (id int, err error) {
	tx, err := tm.db.Begin()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	res, err := tx.Exec("insert into transactions(number,price_total,cost_total,created_at,updated_at) values(?,?,?,?,?)",
		t.Number, t.PriceTotal, t.CostTotal, t.Created_at, t.Updated_at)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	if affected < 1 {
		tx.Rollback()
		return -1, err
	}
	id64, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	for _, v := range d {
		err := tm.addTransactionDetail(tx, v)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
	}
	tx.Commit()
	return int(id64), nil
}

func (tm *TransactionModelImplementation) InsertTransactionID(id int) (err error) {
	tx, err := tm.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec("update transaction_details set transaction_id = ? where transaction_id is null", id)
	if err != nil {
		fmt.Println("error disini")
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
		return err
	}
	tx.Commit()
	return nil
}

func (tm *TransactionModelImplementation) GetTransactionDetailByTransactionID(id int) (details []model.Transaction_Detail, err error) {
	tx, err := tm.db.Begin()
	if err != nil {
		tx.Rollback()
		return details, err
	}
	rows, err := tx.Query("select * from transaction_details where transaction_id = ?", id)
	if err != nil {
		tx.Rollback()
		return details, err
	}
	defer rows.Close()
	for rows.Next() {
		var result model.Transaction_Detail
		var deleted sql.NullTime
		err = rows.Scan(&result.Id, &result.TransactionID, &result.ItemID, &result.ItemQuantity, &result.ItemPrice, &result.ItemCost,
			&result.Created_at, &result.Updated_at, &deleted)
		if err != nil {
			tx.Rollback()
			return details, err
		}
		details = append(details, result)
	}
	tx.Commit()
	return details, nil
}
