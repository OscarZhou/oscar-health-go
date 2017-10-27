package conf

import(
	"github.com/gin-gonic/gin"	
	"github.com/jinzhu/gorm"
	"oscar-health-go/models"
	"oscar-health-go/routers"
	"fmt"
)

var(
	oApp 		*App
)

func init(){
	db := models.NewDatabase()
	router := routers.NewRouter()
	oApp = &App{ Router:router, Db: db}
}

func (a *App) Run(){
	fmt.Println("..................Run App")
	defer a.Db.Close()
	a.Router.Run(":9090")
}

func NewApp() *App{
	fmt.Println("..................New App")
	return oApp
}

type App struct{
	Router 		*gin.Engine
	Db 			*gorm.DB		
}
