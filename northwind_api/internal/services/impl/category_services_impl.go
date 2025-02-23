package impl

import (
	"context"
	"northwind-project/internal/models"
	"northwind-project/internal/repository"
	"northwind-project/internal/services"
)

type CategoryServiceImpl struct {
	repo repository.CategoriesRepository
}

func NewCategoryService(repo repository.CategoriesRepository) services.CategoryService {
	return &CategoryServiceImpl{repo}
}

func (service *CategoryServiceImpl) GetAllCategories(ctx context.Context) ([]models.CategoriesModele, error) {
	return nil, nil
}
func (service *CategoryServiceImpl) GetCategoryById(ctx context.Context, id int) (*models.CategoriesModele, error) {
	return nil, nil
}
func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, category models.CategoriesModele) (*models.CategoriesModele, error) {
	return nil, nil
}
func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, id int, category models.CategoriesModele) (*models.CategoriesModele, error) {
	return nil, nil
}
func (service *CategoryServiceImpl) DeleteCategory(ctx context.Context, id int) error {
	return nil
}
