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
	"strings"
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

type OAuth2 struct{
	Oauth2 		*oauth2.Config
	ThirdParty 	string
	UserInfoURL string
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
// Send the url ("https://accounts.google.com/o/oauth2/auth?
// client_id=841347352749-7htlim54baibrbtet5d3bqc6svrjdgq2.ap
// ps.googleusercontent.com&redirect_uri=http://127.0.0.1:9090/
// auth&response_type=code&scope=https://www.googleapis.com/
// auth/userinfo.email&state=LDO/w1ViP+NvUhGlrk2h29LRIPKY+yHJZB
// QoG0ZqMYo=") the format is like above 
// to server to ask the grant code
func (this *OauthController) LoginByAuth(c *gin.Context) {
	fmt.Println(".......................LoginByAuth")
	// http://localhost:9090?authServer=google/auth0
	// 用这个格式去获取第三方认证服务器，然后获取该认证服务器的相关参数
	authServer := c.DefaultQuery("authServer", "") 
	conf = readCredentialFile("\\conf\\creds.json", authServer)
	// state = randToken()
	// session = sessions.Default(c)
	// session.Set("state", state)
	// session.Save()
	
	state = "xyz"
	redirect := conf.Oauth2.AuthCodeURL(state)
	c.Redirect(http.StatusFound, redirect)
}
// Parse the Credential Information file and fill the OAuth2 
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
// Generate a random value
func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
// Get the token and request the user information from the third authorization server
func (this *OauthController) AuthorizationCodeMethod(c *gin.Context){
	// 查看状态是否有效，实际上就是查看认证服务器返回的状态值与客户端请求授权码时提交的状态值是否一致
	// retriState :=  state
	// queryState := c.Query("state")
	// fmt.Printf("retriState=%v\nqueryState=%v\n", retriState, queryState)
	// if retriState != queryState {
	// 	c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retriState  ))	
	// 	return
	// }

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
	//read the userinfo returned by authorization server and save them into database
	user := &GoogleUser{}
	err = json.Unmarshal(data, user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("user: ", user)
	fmt.Println("Resp body: ", string(data))
	c.Redirect(http.StatusMovedPermanently, "http://localhost:9090/")
}

// User is a retrieved and authentiacted user.
type GoogleUser struct {
    Sub string `json:"sub"`
    Name string `json:"name"`
    GivenName string `json:"given_name"`
    FamilyName string `json:"family_name"`
    Profile string `json:"profile"`
    Picture string `json:"picture"`
    Email string `json:"email"`
    EmailVerified bool `json:"email_verified"`
    Gender string `json:"gender"`
}


func (this *OauthController) TestMethod(c *gin.Context){
	code := c.Query("code")
	url := "https://oscarhealth.au.auth0.com/oauth/token"
	body := "{\"grant_type\":\"authorization_code\",\"client_id\": \"oua-502BEZDq1_2w8wb0h5J-opUOvBy2\",\"client_secret\": \"RdMzoMjq97r596d3IkanEMD28JCZrWa8OpCH0POf5OB0KH8ArosJx4zzgbYlM0td\",\"code\": \""+code+"\",\"redirect_uri\": \"http://localhost:9090/auth\"}"
	payload := strings.NewReader(body)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body1, _ := ioutil.ReadAll(res.Body)

	var token Token
	err := json.Unmarshal(body1, &token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Token:" + token.AccessToken)
	fmt.Println("Token:" + token.TokenID)

	
	// client := conf.Oauth2.Client(oauth2.NoContext, &token)
	// resp, err := client.Get(conf.UserInfoURL)
	// if err != nil{
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// defer resp.Body.Close()
	// data, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println("-------------------")
	// fmt.Println(data)
	sendJWT(token)	

}

func sendJWT(token Token) {
	url := "http://localhost:9090/webapi/v1/brand/4"
	
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("content-type", "application/json")
	authorization := "Bearer "+ token.TokenID
	req.Header.Add("authorization", authorization)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("-------------------")
	fmt.Println(string(body))
	
}

type Token struct{
	AccessToken string `json:"access_token"`
	TokenID string `json:"id_token"`
	ExpiresIn int `json:"expires_in"`
	TokenType string `json:"token_type"`
}