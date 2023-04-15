package main

import (
	"encoding/json"
	"net/http"
	"store/modules/categories"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type product struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name"`
	Price       int                 `json:"price"`
	Category_Id int                 `json:"category_id"`
	Category    categories.Category `json:"category"`
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var product product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db, err := gorm.Open(mysql.Open("root:((root))@tcp(localhost:3306)/store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	result := db.Select("Name", "Price", "Category_Id").Create(&product)
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func getAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db, err := gorm.Open(mysql.Open("root:((root))@tcp(localhost:3306)/store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var products []product
	result := db.Preload("Category").Find(&products)

	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	hasil, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := gorm.Open(mysql.Open("root:((root))@tcp(localhost:3306)/store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var products product
	result := db.Preload("Category").First(&products, id)

	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}

	hasil, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var products product
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	db, errGorm := gorm.Open(mysql.Open("root:((root))@tcp(localhost:3306)/store"), &gorm.Config{})
	if errGorm != nil {
		panic("Failed connect to databases")
	}
	result := db.Model(&product{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":  products.Name,
		"price": products.Price,
	})
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	db, errGorm := gorm.Open(mysql.Open("root:((root))@tcp(localhost:3306)/store"), &gorm.Config{})
	if errGorm != nil {
		panic("Failed connect to databases")
	}
	result := db.Delete(&product{}, id)
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	w.Write([]byte("succes"))
}
