package repository

import (
	"context"
	"northwind-project/internal/entitys"
)

type SuppliersRepository interface {
	GetAll(ctx context.Context) ([]entitys.Categories, error)
	GetByID(ctx context.Context, id int) (entitys.Categories, error)
	Insert(ctx context.Context, categories entitys.Categories) error
	Update(ctx context.Context, categories entitys.Categories) error
	Delete(ctx context.Context, id int) error
}
