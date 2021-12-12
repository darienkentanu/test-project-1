package database

import (
	"database/sql"
	"log"
	"test-project-1/model"
)

type ItemModelImplementation struct {
	db *sql.DB
}

type ItemModel interface {
	AddItem(newItem model.Item) (item model.Item, err error)
	GetAllItems() (items []model.Item, err error)
	DeleteItem(id int) (status string, err error)
	EditItem(newItem model.Item, id int) (item model.Item, err error)
	GetItemById(id int) (item model.Item, err error)
	CheckItemID(id int) (bool, error)
}

func NewItemModel(db *sql.DB) *ItemModelImplementation {
	return &ItemModelImplementation{db: db}
}

func (im *ItemModelImplementation) AddItem(newItem model.Item) (item model.Item, err error) {
	tx, err := im.db.Begin()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	res, err := tx.Exec(
		"insert into items(name, price, cost, created_at, updated_at) values (?,?,?,?,?)",
		newItem.Name, newItem.Price, newItem.Cost, newItem.Created_at, newItem.Updated_at)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	if affected < 1 {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	row := tx.QueryRow("select * from items where id = ?", id)
	err = row.Scan(&item.Id, &item.Name, &item.Price, &item.Cost,
		&item.Created_at, &item.Updated_at)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	tx.Commit()
	return item, err
}

func (im *ItemModelImplementation) GetAllItems() (items []model.Item, err error) {
	tx, err := im.db.Begin()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return items, err
	}
	rows, err := tx.Query("select * from items")
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return items, err
	}
	defer rows.Close()
	for rows.Next() {
		var item = model.Item{}
		err := rows.Scan(&item.Id, &item.Name, &item.Price, &item.Cost, &item.Created_at,
			&item.Updated_at)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return []model.Item{}, err
		}
		items = append(items, item)
	}
	tx.Commit()
	return items, nil
}

func (im *ItemModelImplementation) DeleteItem(id int) (status string, err error) {
	tx, err := im.db.Begin()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	res, err := tx.Exec("delete from items where id = ?", id)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	if affected < 1 {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return "delete success", nil
}

func (im *ItemModelImplementation) EditItem(newItem model.Item, id int) (item model.Item, err error) {
	tx, err := im.db.Begin()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	res, err := tx.Exec(
		"update items set name=?, price=?, cost=?, updated_at=? where id=?",
		newItem.Name, newItem.Price, newItem.Cost, newItem.Updated_at, id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	if affected < 1 {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	row := tx.QueryRow("select * from items where id = ?", id)
	err = row.Scan(&item.Id, &item.Name, &item.Price, &item.Cost,
		&item.Created_at, &item.Updated_at)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	tx.Commit()
	return item, err
}

func (im *ItemModelImplementation) GetItemById(id int) (item model.Item, err error) {
	tx, err := im.db.Begin()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	row := tx.QueryRow("select * from items where id = ?", id)
	err = row.Scan(&item.Id, &item.Name, &item.Price, &item.Cost,
		&item.Created_at, &item.Updated_at)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return item, err
	}
	tx.Commit()
	return item, err
}

func (im *ItemModelImplementation) CheckItemID(id int) (bool, error) {
	tx, err := im.db.Begin()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return false, err
	}
	err = tx.QueryRow("select * from items where id = ?", id).Err()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return false, err
	}
	tx.Commit()
	return true, err
}
