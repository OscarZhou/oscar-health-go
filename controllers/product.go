package controllers

import(
	"github.com/gin-gonic/gin"
	"oscar-go/models"
	"log"
	"net/http"
)

type ProductController struct{
	Controller
}

func (this *ProductController) GetProducts(c *gin.Context){
	sort := c.DefaultQuery("sort", "")
	products, err := models.GetProducts(sort)
	if err != nil{
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, products)
}

