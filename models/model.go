package models

import(
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

const(
	host	= "localhost"
	port	= 5432
	user	= "postgres"
	password= "uiop_098Mn"
	dbname	= "oscarhealth"
)

var(
	db		*gorm.DB
)

func init(){
	fmt.Println("..................Init Database")
	var err error
	db, err = db_new()
	if err != nil{
		fmt.Println("failed to connect database")
	}
	db.AutoMigrate(&Product{}, &Category{}, &Brand{})
}

func NewDatabase() *gorm.DB{
	return db
}


func db_new() (*gorm.DB, error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname= %s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil{
		return &gorm.DB{}, err
	}
	return db, err
}
