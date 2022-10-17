package entity

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Order struct {
	ID           int64     `db:"id"`
	CustomerName string    `db:"customer_name"`
	OrderedAt    time.Time `db:"ordered_at"`
	UpdatedAt    null.Time `db:"updated_at"`
}
