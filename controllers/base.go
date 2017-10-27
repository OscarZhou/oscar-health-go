package controllers

import(
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
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
	Init(ct *gin.Context, cn string)
	Prepare()
	Get(ct *gin.Context)
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Finish()
	Render() error
}

func (c *Controller) Init(ct *gin.Context, cn string){
	c.Data = make(map[interface{}]interface{})
	c.Layout = make([]string, 0)
	c.TplNames = ""
	c.ChildName = cn
	c.Ct = ct
	c.TplExt = "tpl"
}

func (c *Controller) Prepare(){

}

func (c *Controller) Finish(){
	
}

func (ct *Controller) Get(c *gin.Context){
	http.Error(c.Writer, "Method Not Allowed", 405)
}


func (ct *Controller) Post(c *gin.Context){
	http.Error(c.Writer, "Method Not Allowed", 405)
}


func (ct *Controller) Delete(c *gin.Context){
	http.Error(c.Writer, "Method Not Allowed", 405)
}
	
func (ct *Controller) Put(c *gin.Context){
	http.Error(c.Writer, "Method Not Allowed", 405)
}

func (ct *Controller) Head(c *gin.Context){
	http.Error(c.Writer, "Method Not Allowed", 405)
}

func (ct *Controller) Patch(c *gin.Context){
	http.Error(c.Writer, "Method Not Allowed", 405)
}

func (ct *Controller) Options(c *gin.Context){
	http.Error(c.Writer, "Method Not Allowed", 405)
}	

func (ct *Controller) Render() error{
	return nil
}





