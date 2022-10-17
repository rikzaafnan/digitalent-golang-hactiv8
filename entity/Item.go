package entity

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Item struct {
	ID          int64     `db:"id"`
	ItemCode    string    `db:"item_code"`
	Description string    `db:"description"`
	Quantity    int       `db:"quantity"`
	OrderID     string    `db:"order_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   null.Time `db:"updated_at"`
}
