package controllers

import(
	"github.com/gin-gonic/gin"
	"oscar-health-go/models"
	"net/http"
	"strconv"
	"fmt"
)

type BrandController struct{
	Controller
}

func (this *BrandController) GetBrands(c *gin.Context){
	brands := models.GetBrands()
	c.JSON(http.StatusOK, brands)
}

func (this *BrandController) GetBrand(c *gin.Context){
	id := c.Param("id")
	brand := models.GetBrand(id)
	c.JSON(http.StatusOK, brand)
}

func (this *BrandController) CreateBrand(c *gin.Context){
	var brand models.Brand
	err := c.Bind(&brand)
	if err != nil {
		// Redirect to 404
	}
	fmt.Printf("brandname = %s and logo= %s", brand.BrandName, brand.Logo)
	models.CreateBrand(brand)
	c.JSON(http.StatusOK, brand)
}

func (this *BrandController) UpdateBrand(c *gin.Context){
	var newBrand models.Brand
	err := c.Bind(&newBrand)
	if err != nil {
		// Redirect to 404
	}
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("fail to convert the type")
	}
	newBrand.ID = uint(i)
	models.UpdateBrand(newBrand)
	c.JSON(http.StatusOK, newBrand)
}

func (this *BrandController) DeleteBrand(c *gin.Context){
	id := c.Param("id")
	var brand models.Brand
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("fail to convert the type")
	}
	brand.ID = uint(i)
	models.DeleteBrand(brand)
}