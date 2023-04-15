package categories

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetAllCategories() ([]Category, error) {
	var categories []Category
	result := repo.DB.Find(&categories)
	return categories, result.Error
}

func (repo Repository) GetCategoryById(id string) (*Category, error) {
	var category *Category
	result := repo.DB.First(&category, id)
	return category, result.Error
}

func (repo Repository) CreateCategory(category Category) error {
	result := repo.DB.Create(&category)
	return result.Error
}

func (repo Repository) UpdateCategoryById(id string, category Category) error {
	result := repo.DB.Model(&Category{}).Where("id = ?", id).Updates(category)
	return result.Error
}

func (repo Repository) DeleteCategoryById(id string) error {
	result := repo.DB.Delete(&Category{}, id)
	return result.Error
}
