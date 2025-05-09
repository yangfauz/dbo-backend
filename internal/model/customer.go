package model

type Customer struct {
	ID       int    `db:"id" json:"id"`
	Fullname string `db:"fullname" json:"fullname"`
}

func (a *Customer) ToInsert() []interface{} {
	return []interface{}{
		a.Fullname,
	}
}

func (a *Customer) ToUpdate() []interface{} {
	return []interface{}{
		a.ID,
		a.Fullname,
	}
}
