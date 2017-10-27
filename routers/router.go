package routers

import(
	"github.com/gin-gonic/gin"
	"oscar-go/controllers"
	"fmt"
)

var(
	router		*gin.Engine
)

func init(){
	fmt.Println("..................Init Router")
	router = gin.Default()
	configRouter()
}

func NewRouter() *gin.Engine{
	return router
}

func configRouter() {
	fmt.Println("...................Configure Router")
	router.GET("/webapi/v1/product/", (&controllers.ProductController{}).GetProducts)
}
