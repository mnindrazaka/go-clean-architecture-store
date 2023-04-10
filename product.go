package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type product struct {
	Id           int       `json:"id"`
	Nama         string    `json:"nama"`
	Price        int       `json:"price"`
	Categorie_Id int       `json:"categorie_id"`
	categorie    categorie `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var product product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	result := db.Select("Nama", "Price").Create(&product)
	if result.Error != nil {
		w.Write([]byte(result.Error.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func getAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var products []product
	result := db.Find(&products)

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
	db, err := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}
	var products []product
	result := db.First(&products, id)

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
	db, errGorm := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
	if errGorm != nil {
		panic("Failed connect to databases")
	}
	result := db.Model(&product{}).Where("id = ?", id).Updates(map[string]interface{}{
		"nama":  products.Nama,
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
	db, errGorm := gorm.Open(mysql.Open("root:18543@tcp(localhost:3306)/db_store"), &gorm.Config{})
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