package models

import(
	"github.com/jinzhu/gorm"
)

type Category struct{
	gorm.Model
	CategoryName	string
}

func GetProductCategories() []Category {
	var categories []Category
	db.Find(&categories)
	return categories
}

func GetProductCategory(id string) Category {
	var category Category
	db.First(&category, id)
	return category
}

func CreateProductCategory(category Category) {
	db.Create(&category)
}

func UpdateProductCategory(category Category) {
	db.Save(&category)
}

func DeleteProductCategory(category Category) {
	db.Delete(&category)
}
