package order

const (
	FIND_ALL = `
		SELECT o.id, o.order_name, o.customer_id
		FROM orders o
		WHERE 0=0
	`

	FIND_CUSTOMER_ALL = `
		SELECT o.id, o.order_name, o.customer_id, c.fullname as customer_name
		FROM orders o
		join customers c on o.customer_id = c.id
		WHERE 0=0
	`

	FIND_BY_ID = `
		SELECT
			o.id,
			o.order_name,
			o.customer_id
		FROM
			orders o
		WHERE
			o.id = $1;
	`

	INSERT_ORDER = `
	INSERT INTO orders (order_name, customer_id)
	VALUES ($2, $1)
	RETURNING id;
	`

	UPDATE_ORDER = `
	UPDATE orders
	SET order_name = $3, customer_id = $2
	WHERE id = $1;
	`

	DELETE_ORDER = `
	DELETE FROM orders
	WHERE id = $1;
	`
)
