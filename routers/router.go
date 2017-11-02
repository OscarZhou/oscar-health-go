package routers

import(
	"github.com/gin-gonic/gin"
	"oscar-health-go/controllers"
	"fmt"
	"github.com/gin-contrib/sessions"
)

var(
	router		*gin.Engine
	store = sessions.NewCookieStore([]byte("secret"))
)

func init(){
	fmt.Println("..................Init Router")
	router = gin.Default()
	configRouter()
}

func NewRouter() *gin.Engine{
	return router
}

// func loadTemplates(){
// 	var 
// }

func configRouter() {
	fmt.Println("...................Configure Router")
	router.GET("/", (&controllers.DefaultController{}).Index)

	router.Use(sessions.Sessions("gosession", store))
	{
		router.GET("login", (&controllers.OauthController{}).LoginByAuth)
		router.GET("auth", (&controllers.OauthController{}).AuthorizationCodeMethod)
	}


	v1 := router.Group("/webapi/v1", (&controllers.Controller{}).Init)
	{

		v1.GET("/webapi/v1/product/", (&controllers.ProductController{}).GetProducts)
		
		v1.GET("/category/", (&controllers.CategoryController{}).GetProductCategories)
		v1.GET("/category/:id", (&controllers.CategoryController{}).GetProductCategory)
		v1.POST("/category/", (&controllers.CategoryController{}).CreateProductCategory)
		v1.PUT("/category/:id", (&controllers.CategoryController{}).UpdateProductCategory)
		v1.DELETE("/category/:id", (&controllers.CategoryController{}).DeleteProductCategory)		

		
		v1.GET("/brand/", (&controllers.BrandController{}).GetBrands)
		v1.GET("/brand/:id", (&controllers.BrandController{}).GetBrand)
		v1.POST("/brand/", (&controllers.BrandController{}).CreateBrand)
		v1.PUT("/brand/:id", (&controllers.BrandController{}).UpdateBrand)
		v1.DELETE("/brand/:id", (&controllers.BrandController{}).DeleteBrand)		
	}

}
