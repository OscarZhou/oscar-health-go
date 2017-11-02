package controllers

import(
	"golang.org/x/oauth2"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"crypto/rand"	
	"net/http"
)

var (
	conf *OAuth2
	state string
)

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
	UserInfoURL	string
}

type OauthController struct{
	Controller
}

type OAuth2 struct{
	Oauth2 		*oauth2.Config
	ThirdParty 	string
	UserInfoURL string
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
	
	state = randToken()
	// session = sessions.Default(c)
	// session.Set("state", state)
	// session.Save()
	
	redirect := conf.Oauth2.AuthCodeURL(state)
	c.Redirect(http.StatusFound, redirect)
}

func readCredentialFile(fileName string, authServer string) *OAuth2 {
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
			return &OAuth2{
				Oauth2: &oauth2.Config {
					ClientID:		v.Cid,
					ClientSecret:	v.Csecret,
					RedirectURL:	v.RedirectURL,
					Scopes:			v.Scopes,
					Endpoint: oauth2.Endpoint{
						AuthURL:	v.AuthURL,
						TokenURL:	v.TokenURL,
					},
				},
				ThirdParty:		v.ThirdParty,
				UserInfoURL:	v.UserInfoURL,
			}
		}
	}
	return &OAuth2{}
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func (this *OauthController) AuthorizationCodeMethod(c *gin.Context){
	// 查看状态是否有效，实际上就是查看认证服务器返回的状态值与客户端请求授权码时提交的状态值是否一致
	retriState :=  state
	queryState := c.Query("state")
	fmt.Printf("retriState=%v\nqueryState=%v\n", retriState, queryState)
	if retriState != queryState {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retriState  ))	
		return
	}

	// c.Query("code") 得到的值是认证服务器返回的授权码
	// Exchange 是利用授权码为了向服务器申请令牌. Exchange方法调用分布在多个文件中
	// retrieveToken(golang.org/x/oauth2/oauth.go)-> 
	// internal.RetrieveToken (golang.org/x/oauth2/token.go)->
	// RetrieveToken (golang.org/x/oauth2/internal/token.go)
	tok, err := conf.Oauth2.Exchange(oauth2.NoContext, c.Query("code")) 
	if err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	client := conf.Oauth2.Client(oauth2.NoContext, tok)
	resp, err := client.Get(conf.UserInfoURL)
	if err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Resp body: ", string(data))
	

	c.Redirect(http.StatusMovedPermanently, "http://localhost:9090/")

	//pwd, _ := os.Getwd()
	//redirect := pwd + "\\views\\index.html"
	// fmt.Printf("..............redirect=%v\n", redirect)
	// c.HTML(http.StatusOK, "Home", redirect)
	//c.Redirect(http.StatusFound, redirect)
	
	//t, _ := template.ParseFiles(redirect)
	//t.Execute(c.Writer, nil)

}