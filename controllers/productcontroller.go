package controllers

import(
	"github.com/gin-gonic/gin"
	"oscar-health-go/models"
	"net/http"
)

type ProductController struct{
	Controller
}

func (this *ProductController) GetProducts(c *gin.Context){
	sort := c.DefaultQuery("sort", "")
	products := models.GetProducts(sort)
	
	c.JSON(http.StatusOK, products)
}

func (this *ProductController) GetProduct(c *gin.Context){
	id := c.Param("id")
	product := models.GetProduct(id)

	c.JSON(http.StatusOK, product)
}