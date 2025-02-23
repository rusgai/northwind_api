package services

import (
	"context"
	"northwind-project/internal/models"
)

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]models.CategoriesModele, error)
	GetCategoryById(ctx context.Context, id int) (*models.CategoriesModele, error)
	CreateCategory(ctx context.Context, category models.CategoriesModele) (*models.CategoriesModele, error)
	UpdateCategory(ctx context.Context, id int, category models.CategoriesModele) (*models.CategoriesModele, error)
	DeleteCategory(ctx context.Context, id int) error
}
