package categories

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.Usecase.CreateCategory(category)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func (handler Handler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	categories, err := handler.Usecase.GetAllCategories()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	hasil, err := json.Marshal(categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func (handler Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	category, err := handler.Usecase.GetCategoryById(id)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	hasil, err := json.Marshal(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(hasil)
}

func (handler Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = handler.Usecase.UpdateCategoryById(id, category)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("succes"))
}

func (handler Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	err := handler.Usecase.DeleteCategoryById(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("succes"))
}
