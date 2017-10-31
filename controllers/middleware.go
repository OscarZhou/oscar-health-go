package controllers

import(
	"github.com/gin-gonic/gin"
	"fmt"
)

type MiddlewareController struct{
	Controller
}

func (this *MiddlewareController) Cors(c *gin.Context) {
	fmt.Println("....................Prepare")
	c.Next()



	// this.Ct.Set("Access-Control-Allow-Origin", "*")
	// this.Ct.Set("Access-Control-Allow-Headers", "Content-Type")
	// this.Ct.Set("Access-Control-Allow-Methods", "GET, POST, PUT")
	// this.Ct.Next()
	// fmt.Println(this.Ct.Get("Access-Control-Allow-Origin"))
	// fmt.Println(this.Ct.Get("Access-Control-Allow-Headers"))
	// fmt.Println(this.Ct.Get("Access-Control-Allow-Methods"))
}
