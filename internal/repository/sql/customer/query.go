package customer

const (
	FIND_ALL = `
		SELECT c.id, c.fullname
		FROM customers c
		WHERE 0=0
	`

	FIND_BY_ID = `
		SELECT 
			c.id, 
			c.fullname
		FROM 
			customers c
		WHERE 
			c.id = $1;
	`
	INSERT_CUSTOMER = `
	INSERT INTO customers (fullname)
	VALUES ($1)
	RETURNING id;
	`
	UPDATE_CUSTOMER = `
	UPDATE customers
	SET fullname = $2
	WHERE id = $1;
	`
	DELETE_CUSTOMER = `
	DELETE FROM customers
	WHERE id = $1;
	`
)
