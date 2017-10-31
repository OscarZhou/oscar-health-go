package controllers

import(
	"golang.org/x/oauth2"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"github.com/gin-contrib/sessions"
	"crypto/rand"	
	"net/http"
)

var conf *oauth2.Config
// var store = sessions.NewCookieStore([]byte("secret"))

type Creds struct{
	Creds []Credential 
}

type Credential struct {
	ThirdParty	string
	Cid			string `json:"cid"`
	Csecret		string `json:"csecret"`
	RedirectURL	string
	Scopes		[]string
	AuthURL		string
	TokenURL	string
}

type OauthController struct{
	Controller
}


func (this *OauthController) Prepare(c *gin.Context) {
	fmt.Println("...................OauthController Prepare Cors")
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT")
	
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
	c.Next()
}

func (this *OauthController) LoginByAuth(c *gin.Context) {
	fmt.Println(".......................LoginByAuth")
	authServer := c.DefaultQuery("authServer", "")
	conf = readCredentialFile("\\conf\\creds.json", authServer)
	state := randToken()
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()

	redirect := conf.AuthCodeURL(state)
	c.Redirect(http.StatusFound, redirect)
}

func readCredentialFile(fileName string, authServer string) *oauth2.Config {
	creds := &Creds{}
	pwd, _ := os.Getwd()
	file, err := ioutil.ReadFile(pwd + fileName)
	if err != nil {
		fmt.Printf("...............File error: %v\n", err)
		os.Exit(1)
	}
	err = json.Unmarshal(file, creds)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range creds.Creds{
		if v.ThirdParty == authServer {
			return &oauth2.Config{
				ClientID:		v.Cid,
				ClientSecret:	v.Csecret,
				RedirectURL:	v.RedirectURL,
				Scopes:			v.Scopes,
				Endpoint: oauth2.Endpoint{
					AuthURL:	v.AuthURL,
					TokenURL:	v.TokenURL,
				},
			}
		}
	}
	return &oauth2.Config{}
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func getLoginURL(state string) string{
	return conf.AuthCodeURL(state)
}
