package controllers

import(
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
)

type DefaultController struct{
	Controller
}

func (this *DefaultController) Index(c *gin.Context) {
	this.Ct = c
	this.TplNames = "\\views\\index.html"
	pwd, _ := os.Getwd()
	
	t, _ := template.ParseFiles(pwd + this.TplNames)
	t.Execute(c.Writer, nil)
}


