package main

import (
	"log"
	"net/http"
	"northwind-project/internal/configuration"
	"northwind-project/internal/controllers"
	"northwind-project/internal/repository/pgimpl"
	"northwind-project/internal/services/impl"
	"os"
)

func run() error {
	router := http.NewServeMux()
	loger := log.New(os.Stdout, "northwind-api", log.LstdFlags|log.Lshortfile)
	conf := configuration.NewConfig(".env")
	db := configuration.NewPgDb(conf)
	categoryRepository := pgimpl.NewCategoriesRepositoryPgImpl(db)
	categoryService := impl.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService, loger)
	categoryController.RegisterRoutes(router)
	server := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}
	loger.Println("Server is running on port 8001")
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
