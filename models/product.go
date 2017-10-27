package models

import(
	"github.com/jinzhu/gorm"
	"fmt"
)

type Product struct{
	gorm.Model
	CategoryID 			int			
	BrandID				int				
	ProductName 		string				
	RegularPrice 		float64	
	SpecialPrice 		float64
	NZRegularPrice 		float64
	NZSpecialPrice 		float64
	Image 				string
	Display 			string
	Description 		string	
}

func GetProducts(sort string) []Product {
	var products []Product

	if sort == ""{
		fmt.Println("db.Find")
		db.Find(&products)
	}else{
		if sort[0] == '-'{
			if sort[1:] == "category"{
				db.Order("category_id DESC").Find(&products)
			}
		}else {
			if sort == "category"{
				db.Order("category_id ASC").Find(&products)
			}
		}
	}

	return products
}

func GetProduct(id string) Product {
	var product Product
	db.First(&product, id)
	return product
}