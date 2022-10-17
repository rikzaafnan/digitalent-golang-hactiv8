package item_pg

import (
	"assignment-2/dto"
	"assignment-2/entity"
	itemrepository "assignment-2/repository/item_repository"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	sqlItem = `
		SELECT i.id, i.item_code,i.description, i.quantity, i.order_id, i.created_at, i.updated_at 
		from items as i`
	sqlItemByOrderId = `SELECT i.id, i.item_code,i.description, i.quantity, i.order_id, i.created_at, i.updated_at 
	from items as i where i.order_id= $1`
	sqlInsertItem = `INSERT INTO items
	(
		item_code, 
		description,
		quantity,
		order_id,
		created_at
	)
	VALUES($1, $2, $3, $4, $5) RETURNING id;`

	readItem = `SELECT i.id, i.item_code,i.description, i.quantity, i.order_id, i.created_at, i.updated_at 
	from items as i
		where i.id = $1`

	updateItem = `UPDATE items
	SET item_code = $2,
	description = $3,
	quantity = $4,
	order_id = $5,
	updated_at = $6,
	WHERE id = $1`

	deleteItem = `DELETE FROM items
WHERE id = $1;`
)

type itemPG struct {
	db *sqlx.DB
}

func NewItemPG(db *sqlx.DB) itemrepository.ItemRepository {
	return &itemPG{
		db: db,
	}
}

func (i *itemPG) CreateItem(req *dto.ItemRequest) (int64, error) {

	var id int
	tx := i.db.MustBegin()
	err := tx.QueryRowx(sqlInsertItem, req.ItemCode, req.Description, req.Quantity, req.OrderId, time.Now()).Scan(&id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if id == 0 {
		// log.Println(err)
		return 0, errors.New("not insert to database")
	}

	return int64(id), nil

}

func (i *itemPG) FindItemById(itemId int64) (entity.Item, error) {

	var item entity.Item

	err := i.db.Get(&item, sqlItem+" where i.id =? limit = 1 ", itemId)
	if err != nil {

		log.Println(sqlItem+" where i.id =? limit = 1 ", itemId)

		return item, err
	}

	return item, nil

}

func (i *itemPG) FindAllItemByOrderID(orderID int64) ([]entity.Item, error) {

	var items []entity.Item

	err := i.db.Select(&items, sqlItemByOrderId, orderID)
	if err != nil {
		log.Println(sqlItem+" where i.order_id =?", orderID)
		return nil, err
	}

	return items, nil

}

func (i *itemPG) UpdateItem(itemId int64, req *dto.ItemRequest) (int64, error) {

	result, err := i.db.Exec(updateItem, itemId, req.ItemCode, req.Description, req.Quantity, req.OrderId, time.Now())
	if err != nil {
		return 0, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowAffected, nil

}

func (i *itemPG) DeleteItem(itemId int64) (bool, error) {

	result, err := i.db.Exec(deleteItem, itemId)
	if err != nil {
		return false, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return false, err
	}

	return true, nil

}
