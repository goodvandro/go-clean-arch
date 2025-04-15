package database

import (
	"database/sql"

	"github.com/goodvandro/go-clean-arch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	query := "INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)"
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
