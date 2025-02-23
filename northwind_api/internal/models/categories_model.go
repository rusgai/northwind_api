package models

type CategoriesModele struct {
	ID            int    `json:"category_id"`
	Category_name string `json:"category_name"`
	Description   string `json:"description"`
	Picture       []byte `json:"picture"`
}
