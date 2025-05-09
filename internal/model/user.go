package model

type User struct {
	ID       int    `db:"id" json:"id"`
	Fullname string `db:"fullname" json:"fullname"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

func (a *User) ToInsert() []interface{} {
	return []interface{}{
		a.Email,
		a.Password,
		a.Fullname,
	}
}
