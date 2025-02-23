package entitys

type Categories struct {
	ID            int    `db:"category_id"`
	Category_name string `db:"category_name"`
	Description   string `db:"description"`
	Picture       []byte `db:"picture"`
}
