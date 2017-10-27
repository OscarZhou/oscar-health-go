package routers

import(
	"github.com/gin-gonic/gin"
	"oscar-health-go/controllers"
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

	router.GET("/webapi/v1/category/", (&controllers.CategoryController{}).GetProductCategories)
	router.GET("/webapi/v1/category/:id", (&controllers.CategoryController{}).GetProductCategory)
	router.POST("/webapi/v1/category/", (&controllers.CategoryController{}).CreateProductCategory)
	router.PUT("/webapi/v1/category/:id", (&controllers.CategoryController{}).UpdateProductCategory)
	router.DELETE("/webapi/v1/category/:id", (&controllers.CategoryController{}).DeleteProductCategory)
}
