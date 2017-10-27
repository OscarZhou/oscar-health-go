package models

import(
	// "database/sql"				//database plugin
	
	"time"
)

type Product struct{
	Id 					int			
	CategoryId 			int			
	BrandId				int				
	Name 				string				
	RegularPrice 		float64	
	SpecialPrice 		float64
	NZRegularPrice 		float64
	NZSpecialPrice 		float64
	Image 				string
	Display 			string
	Description 		string		
	Created				time.Time
}

func GetProducts(sort string) ([]Product, error){
	var orderby string
	if sort == ""{
		orderby = ""
	}else{
		if sort[0] == '-'{
			if sort[1:] == "category"{
				orderby = "ORDER BY category_id DESC"
			}
		}else {
			if sort == "category"{
				orderby = "ORDER BY category_id ASC"
			}
		}
	}

	sqlStatement := `
	SELECT product_id, category_id, brand_id, product_name, regular_price,
	special_price, nz_regular_price, nz_special_price, image, display, description
	 FROM oh_product ` + orderby + ` LIMIT 10;`

	var products []Product
	rows, err := db.Query(sqlStatement)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	var product Product
	for rows.Next(){
		err = rows.Scan(&product.Id,&product.CategoryId,&product.BrandId,
			&product.Name, &product.RegularPrice, &product.SpecialPrice,
			&product.NZRegularPrice, &product.NZSpecialPrice, &product.Image,
			&product.Display, &product.Description)
		if err != nil{
			panic(err)
		}
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil{
		panic(err)
	}
	
	return products, nil
}
