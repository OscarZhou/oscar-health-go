package controllers

import(
	"github.com/gin-gonic/gin"
	"oscar-health-go/models"
	"net/http"
	"strconv"
	"fmt"
)

type CategoryController struct{
	Controller
}

func (this *CategoryController) GetProductCategories(c *gin.Context){
	fmt.Println("........................GetProductCategories")
	categories := models.GetProductCategories()
	c.JSON(http.StatusOK, categories)
}

func (this *CategoryController) GetProductCategory(c *gin.Context){
	id := c.Param("id")
	category := models.GetProductCategory(id)
	c.JSON(http.StatusOK, category)
}

func (this *CategoryController) CreateProductCategory(c *gin.Context){
	var category models.Category
	err := c.Bind(&category)
	if err != nil {
		// Redirect to 404
	}
	models.CreateProductCategory(category)
	c.JSON(http.StatusOK, category)
}

func (this *CategoryController) UpdateProductCategory(c *gin.Context){
	var newCategory models.Category
	err := c.Bind(&newCategory)
	if err != nil {
		// Redirect to 404
	}
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("fail to convert the type")
	}
	newCategory.ID = uint(i)
	models.UpdateProductCategory(newCategory)
	c.JSON(http.StatusOK, newCategory)
}

func (this *CategoryController) DeleteProductCategory(c *gin.Context){
	id := c.Param("id")
	var category models.Category
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("fail to convert the type")
	}
	category.ID = uint(i)
	models.DeleteProductCategory(category)
}