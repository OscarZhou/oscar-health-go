package controllers

import(
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"fmt"
)

type Controller struct{
	Ct 			*gin.Context
	Tpl			*template.Template
	Data		map[interface{}]interface{}
	ChildName	string
	TplNames		string
	Layout		[]string
	TplExt		string
}

type ControllerInterface interface{
	Init(c *gin.Context)
	Prepare(c *gin.Context)
	Get(ct *gin.Context)
	Post(c *gin.Context)
	Delete(c *gin.Context)
	Put(c *gin.Context)
	Head(c *gin.Context)
	Patch(c *gin.Context)
	Options(c *gin.Context)
	Finish()
	Render() error
}

func (ct *Controller) Init(c *gin.Context){
	fmt.Println(".....................Init Base Controller")
	ct.Data = make(map[interface{}]interface{})
	ct.Layout = make([]string, 0)
	ct.TplNames = ""
	ct.ChildName = ""
	ct.Ct = c
	ct.TplExt = "tpl"
}

func (ct *Controller) Prepare(c *gin.Context){
	http.Error(ct.Ct.Writer, "Method Not Allowed", 405)
}

func (c *Controller) Finish(){
	
}

// func (ct *Controller) Get(){
// 	http.Error(ct.Ct.Writer, "Method Not Allowed", 405)
// }

func (ct *Controller) Get(c *gin.Context){
	http.Error(ct.Ct.Writer, "Method Not Allowed", 405)
}


func (ct *Controller) Post(c *gin.Context){
	http.Error(ct.Ct.Writer, "Method Not Allowed", 405)
}


func (ct *Controller) Delete(c *gin.Context){
	http.Error(ct.Ct.Writer, "Method Not Allowed", 405)
}
	
func (ct *Controller) Put(c *gin.Context){
	http.Error(ct.Ct.Writer, "Method Not Allowed", 405)
}

func (ct *Controller) Head(c *gin.Context){
	http.Error(ct.Ct.Writer, "Method Not Allowed", 405)
}

func (ct *Controller) Patch(c *gin.Context){
	http.Error(ct.Ct.Writer, "Method Not Allowed", 405)
}

func (ct *Controller) Options(c *gin.Context){
	http.Error(ct.Ct.Writer, "Method Not Allowed", 405)
}	

func (ct *Controller) Render() error{
	return nil
}




