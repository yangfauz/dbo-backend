package model

type Order struct {
	ID        int    `db:"id" json:"id"`
	CusomerID int    `db:"customer_id" json:"customer_id"`
	OrderName string `db:"order_name" json:"order_name"`
}

type OrderCustomer struct {
	ID           int    `db:"id" json:"id"`
	OrderName    string `db:"order_name" json:"order_name"`
	CusomerID    int    `db:"customer_id" json:"customer_id"`
	CustomerName string `db:"customer_name" json:"customer_name"`
}

func (a *Order) ToInsert() []interface{} {
	return []interface{}{
		a.CusomerID,
		a.OrderName,
	}
}
func (a *Order) ToUpdate() []interface{} {
	return []interface{}{
		a.ID,
		a.CusomerID,
		a.OrderName,
	}
}
