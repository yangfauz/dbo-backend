package user

const (
	FIND_BY_ID = `
	SELECT 
		u.id, 
		u.email,  
		u.password, 
		u.fullname
	FROM 
		users u
	WHERE 
		u.id = $1;
`

	FIND_BY_EMAIL = `
	SELECT 
		u.id, 
		u.email,  
		u.password, 
		u.fullname
	FROM 
		users u
	WHERE 
		u.email = $1;
`
	INSERT_USER = `
	INSERT INTO users (email, password, fullname)
	VALUES ($1, $2, $3)
	RETURNING id;
`
)
