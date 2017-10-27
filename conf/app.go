package conf

import(
	"github.com/gin-gonic/gin"	
	"database/sql"				//database plugin
	"oscar-go/models"
	"oscar-go/routers"
	"fmt"
)

var(
	OApp 		*App
)

func init(){
	db := models.NewDatabase()
	router := routers.NewRouter()
	OApp = &App{ Router:router, Db: db}
}

func (a *App) Run(){
	fmt.Println("..................Run App")
	defer a.Db.Close()
	a.Router.Run(":9090")
}

func NewApp() *App{
	fmt.Println("..................New App")
	return OApp
}

type App struct{
	Router 		*gin.Engine
	Db 			*sql.DB		
}
