package pgimpl

import (
	"context"
	"northwind-project/internal/entitys"
	"northwind-project/internal/repository"

	"github.com/jmoiron/sqlx"
)

type CategoriesRepositoryPgImpl struct {
	db *sqlx.DB
}

func NewCategoriesRepositoryPgImpl(db *sqlx.DB) repository.CategoriesRepository {
	return &CategoriesRepositoryPgImpl{db}
}

func (repo *CategoriesRepositoryPgImpl) GetAll(ctx context.Context) ([]entitys.Categories, error) {

	qury := `SELECT * FROM categories`
	var categories []entitys.Categories
	err := repo.db.SelectContext(ctx, &categories, qury)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (repo *CategoriesRepositoryPgImpl) GetByID(ctx context.Context, id int) (*entitys.Categories, error) {
	return nil, nil
}
func (repo *CategoriesRepositoryPgImpl) Insert(ctx context.Context, categories entitys.Categories) error {
	return nil
}
func (repo *CategoriesRepositoryPgImpl) Update(ctx context.Context, categories entitys.Categories) error {
	return nil
}
func (repo *CategoriesRepositoryPgImpl) Delete(ctx context.Context, id int) error {
	return nil
}
