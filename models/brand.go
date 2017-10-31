package models

import(
	"github.com/jinzhu/gorm"
)

type Brand struct{
	gorm.Model
	BrandName	string
	Logo		string
}

func GetBrands() []Brand {
	var brands []Brand
	db.Find(&brands)
	return brands
}

func GetBrand(id string) Brand {
	var brand Brand
	db.First(&brand, id)
	return brand
}

func CreateBrand(brand Brand) {
	db.Create(&brand)
}

func UpdateBrand(brand Brand) {
	db.Save(&brand)
}

func DeleteBrand(brand Brand) {
	db.Delete(&brand)
}
