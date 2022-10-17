package orderpg

import (
	"assignment-2/dto"
	"assignment-2/entity"
	"assignment-2/repository/orderrepository"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	sqlOrder = `
		SELECT o.id, o.customer_name, o.ordered_at, o.updated_at
		from orders as o
		ORDER BY o.id ASC;`

	sqlCreateOrder = `INSERT INTO orders
	(
		customer_name, 
		ordered_at
	)
	VALUES($1, $2) RETURNING id;`

	readOrder = `SELECT o.id, o.customer_name, o.ordered_at, o.updated_at 
	from orders as o
		LEFT JOIN items as i ON i.order_id = o.id
		where o.id = $1
	ORDER BY o.id ASC;`

	updateOrder = `UPDATE orders
	SET customer_name = $2,
	ordered_at = $3
	WHERE id = $1`

	deleteOrder = `DELETE FROM orders
WHERE id = $1;`

	sqlInsertItem = `INSERT INTO items
	(
		item_code, 
		description,
		quantity,
		order_id,
		created_at
	)
	VALUES($1, $2, $3, $4, $5) RETURNING id;`

	sqlUpdateItem = `UPDATE items
	SET item_code = $2,
	description = $3,
	quantity = $4,
	order_id = $5,
	updated_at = $6
	WHERE id = $1`
)

type orderPG struct {
	db *sqlx.DB
}

func NewOrderPG(db *sqlx.DB) orderrepository.OrderRepository {
	return &orderPG{
		db: db,
	}
}

func (o *orderPG) CreateOrder(req *dto.OrderRequest) (int64, error) {

	var id int
	tx := o.db.MustBegin()
	err := tx.QueryRowx(sqlCreateOrder, req.CustomerName, req.OrderedAt).Scan(&id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return 0, err
	}

	if len(req.OrderItems) > 0 {

		orderId := id
		for _, item := range req.OrderItems {

			itemRequest := dto.ItemRequest{}
			itemRequest.ItemCode = item.ItemCode
			itemRequest.Description = item.Description
			itemRequest.Quantity = item.Quantity
			itemRequest.OrderId = orderId

			_, err := tx.Exec(sqlInsertItem, itemRequest.ItemCode, itemRequest.Description, itemRequest.Quantity, itemRequest.OrderId, time.Now())
			if err != nil {
				tx.Rollback()
				log.Println(err)
				return 0, err
			}

		}

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

func (o *orderPG) FindOrderById(orderId int64) (entity.Order, error) {

	var p entity.Order

	err := o.db.Get(&p, readOrder, orderId)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (o *orderPG) FindAllOrder() ([]entity.Order, error) {

	var orders []entity.Order

	err := o.db.Select(&orders, sqlOrder)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderPG) UpdateOrder(orderId int64, req *dto.OrderRequestUpdate) (int64, error) {

	tx := o.db.MustBegin()
	result, err := tx.Exec(updateOrder, orderId, req.CustomerName, req.OrderedAt)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return 0, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if len(req.OrderItems) > 0 {

		for _, item := range req.OrderItems {

			itemRequest := dto.ItemRequest{}
			itemRequest.ItemCode = item.ItemCode
			itemRequest.Description = item.Description
			itemRequest.Quantity = item.Quantity
			itemRequest.OrderId = int(orderId)

			_, err := tx.Exec(sqlUpdateItem, item.ID, itemRequest.ItemCode, itemRequest.Description, itemRequest.Quantity, itemRequest.OrderId, time.Now())
			if err != nil {
				tx.Rollback()
				log.Println(err)
				return 0, err
			}

		}

	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return rowAffected, nil
}

func (o *orderPG) DeleteOrder(orderId int64) (bool, error) {

	// result, err := o.db.Exec(deleteOrder, orderId)
	// if err != nil {
	// 	return false, err
	// }
	// _, err = result.RowsAffected()
	// if err != nil {
	// 	return false, err
	// }

	tx := o.db.MustBegin()

	_, err := tx.Exec("DELETE FROM items WHERE order_id = $1", orderId)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return false, err
	}

	result, err := tx.Exec(deleteOrder, orderId)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return false, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}
