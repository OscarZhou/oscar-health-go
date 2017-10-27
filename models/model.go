package models

import(
	"database/sql"				//database plugin
	_ "github.com/lib/pq"		//postgresql plugin, which needs to use together with "database/sql"
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
	db		*sql.DB
)

func init(){
	fmt.Println("..................Init Database")
	var err error
	db, err = db_new()
	if err != nil{
		fmt.Println("连接数据库失败")
	}
}

func NewDatabase() *sql.DB{
	return db
}


func db_new() (*sql.DB, error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		return &sql.DB{}, err
	}

	err = db.Ping()
	if err != nil{
		return &sql.DB{}, err
	}
	return db, nil
}
