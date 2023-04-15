package main

import (
	"fmt"
	"net/http"
	"store/modules/categories"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:((root))@tcp(localhost:3306)/store"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	categoryRepo := categories.Repository{DB: db}
	categoryUsecase := categories.Usecase{Repo: categoryRepo}
	categoryHandler := categories.Handler{Usecase: categoryUsecase}

	r := mux.NewRouter()
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/products", jwtMiddleware(createProduct)).Methods("POST")
	r.HandleFunc("/products", jwtMiddleware(getAllProduct)).Methods("GET")
	r.HandleFunc("/products/{id}", jwtMiddleware(getProduct)).Methods("GET")
	r.HandleFunc("/products/{id}", jwtMiddleware(updateProduct)).Methods("PUT")
	r.HandleFunc("/products/{id}", jwtMiddleware(deleteProduct)).Methods("DELETE")

	r.HandleFunc("/categories", jwtMiddleware(categoryHandler.CreateCategory)).Methods("POST")
	r.HandleFunc("/categories", jwtMiddleware(categoryHandler.GetAllCategories)).Methods("GET")
	r.HandleFunc("/categories/{id}", jwtMiddleware(categoryHandler.GetCategory)).Methods("GET")
	r.HandleFunc("/categories/{id}", jwtMiddleware(categoryHandler.UpdateCategory)).Methods("PUT")
	r.HandleFunc("/categories/{id}", jwtMiddleware(categoryHandler.DeleteCategory)).Methods("DELETE")

	PORT := 3000
	fmt.Println("starting web server at localhost:", PORT)
	if err := http.ListenAndServe(":"+strconv.Itoa(PORT), r); err != nil {
		fmt.Println(err)
		return
	}
}
