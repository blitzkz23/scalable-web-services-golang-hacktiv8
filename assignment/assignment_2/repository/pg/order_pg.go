package pg

import (
	"database/sql"
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/entity"
	"scalable-web-services-golang-hacktiv8/assignment/assignment_2/repository"
)

// ! Query
const (
	insertOrderQuery = `
		INSERT into orders (customer_name)
		VALUES ($1)
		RETURNING id, customer_name, ordered_at;
	`
	retrieveOrderQuery = `
		SELECT id, customer_name, ordered_at
		FROM orders
	`
)

// ! Repository implementation dari order
type orderPg struct {
	db *sql.DB
}

// * Factory function yang mengembalikan orderPg dengan inject db.
func NewOrderPg(db *sql.DB) repository.OrderRepository {
	return &orderPg{
		db: db,
	}
}

// * Implement func from interface
func (o *orderPg) InsertOrder(orderPayload *entity.Order) (*entity.Order, error) {
	// ! Insert data order ke database
	var order entity.Order

	row := o.db.QueryRow(insertOrderQuery, orderPayload.Customer_name)
	err := row.Scan(&order.ID, &order.Customer_name, &order.Ordered_at)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *orderPg) GetAllOrders() ([]*entity.Order, error) {
	// ! Ambil data order dari database
	var orders []*entity.Order
	rows, err := o.db.Query(retrieveOrderQuery)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order entity.Order
		err := rows.Scan(&order.ID, &order.Customer_name, &order.Ordered_at)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}
