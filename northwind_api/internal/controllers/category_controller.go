package controllers

import (
	"log"
	"net/http"
	"northwind-project/internal/services"
)

type CategoryController struct {
	sevice services.CategoryService
	loger  *log.Logger
}

func NewCategoryController(s services.CategoryService, logger *log.Logger) *CategoryController {
	return &CategoryController{}
}
func (controller *CategoryController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/categories", controller.getAllCategories)
	mux.HandleFunc("/api/v1/categories/", controller.getCategoryById)
	mux.HandleFunc("/api/v1/categories", controller.createCategory)
	mux.HandleFunc("/api/v1/categories/", controller.updateCategory)
	mux.HandleFunc("/api/v1/categories/", controller.deleteCategory)
}
func (controller *CategoryController) getAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllCategories"))
}
func (controller *CategoryController) getCategoryById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetCategoryById"))
}
func (controller *CategoryController) createCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateCategory"))
}
func (controller *CategoryController) updateCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateCategory"))
}
func (controller *CategoryController) deleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteCategory"))
}
