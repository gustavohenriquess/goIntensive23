package database

import (
	"database/sql"

	"github.com/gustavohenriquess/go-intensive23/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) (int64, error) {
	result, err := r.Db.Exec("INSERT INTO orders (price, tax, final_price) VALUES (?, ?, ?)",
		order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OrderRepository) GetTotalTransactions() (int, error) {
	var total int

	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
