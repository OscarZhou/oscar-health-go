package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	// "html/template"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	// "github.com/auth0/go-jwt-middleware"
	// "github.com/dgrijalva/jwt-go"
	auth0 "github.com/auth0-community/go-auth0"
	jose "gopkg.in/square/go-jose.v2"
	jwt "gopkg.in/square/go-jose.v2/jwt"
)

var (
	conf  *OAuth2
	state string
)

type Creds struct {
	Creds []Credential
}

type Credential struct {
	ThirdParty  string
	Cid         string `json:"cid"`
	Csecret     string `json:"csecret"`
	RedirectURL string
	Scopes      []string
	AuthURL     string
	TokenURL    string
	UserInfoURL string
}

type OAuth2 struct {
	Oauth2      *oauth2.Config
	ThirdParty  string
	UserInfoURL string
}

type OauthController struct {
	Controller
}

func (oauthController *OauthController) Prepare(c *gin.Context) {
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
func (oauthController *OauthController) LoginByAuth(c *gin.Context) {
	fmt.Println(".......................LoginByAuth")

	// http://localhost:9090?authServer=google/auth0
	// 用这个格式去获取第三方认证服务器，然后获取该认证服务器的相关参数
	authServer := c.DefaultQuery("authServer", "")
	conf = readCredentialFile("\\conf\\creds.json", authServer)

	state := "xyz"
	redirectURI := conf.Oauth2.AuthCodeURL(state)
	c.Redirect(http.StatusFound, redirectURI)
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
	for _, v := range creds.Creds {
		if v.ThirdParty == authServer {
			return &OAuth2{
				Oauth2: &oauth2.Config{
					ClientID:     v.Cid,
					ClientSecret: v.Csecret,
					RedirectURL:  v.RedirectURL,
					Scopes:       v.Scopes,
					Endpoint: oauth2.Endpoint{
						AuthURL:  v.AuthURL,
						TokenURL: v.TokenURL,
					},
				},
				ThirdParty:  v.ThirdParty,
				UserInfoURL: v.UserInfoURL,
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

var session map[string]string

// Get the token and request the user information from the third authorization server
func (oauthController *OauthController) AuthorizationCodeMethod(c *gin.Context) {
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

	// tok, err := conf.Oauth2.Exchange(oauth2.NoContext, c.Query("code"))
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }
	session = make(map[string]string, 2)
	code := c.Query("code")
	frontToken := requestToken(code, "authorization_code")
	session["front_token"] = frontToken.AccessToken
	fmt.Printf("================session token:%v\n", session["front_token"])
	// client := conf.Oauth2.Client(oauth2.NoContext, tok)
	// resp, err := client.Get(conf.UserInfoURL)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// c.Redirect(http.StatusMovedPermanently, "http://localhost:9090/")
}

// 获取Auth0返回的grant code，然后请求access token，并且模拟发送认证请求头
// "Authorization": "Bearer *token*"
// this token is probably JWT
func EmulateFrontEndSendToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("...............EmulateFrontEndSendToken")
		//code := c.Query("code")
		//前端authorization_code 模式去请求token
		//该参数的token 包括access_token, token_id, expires_in, token_type
		//frontToken := requestToken(code, "authorization_code")
		frontToken := session["front_token"]
		// fmt.Printf("================成功获取客户端token:%v\n", frontToken.AccessToken)
		fmt.Printf("================成功获取客户端token:%v\n", frontToken)
		//模拟请求头
		c.Request.Header.Set("Authorization", "Bearer "+frontToken)
		c.Next()
	}
}

var (
	clientID     = "5dmeY74FoZ4Ua39jpp7fK8IDCthZwgMT"
	clientSecret = "XFq4yszVca4BfR_ZXHtc-6FeikjTOwXiBVRJ_m7X9JOCU11gVVbdwus4HMIgQyyr"
	audience     = "http://localhost:9090/brand"
	publicKey	= `
	MIIDCzCCAfOgAwIBAgIJN5dLzmhHYFetMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNV
	BAMTGG9zY2FyaGVhbHRoLmF1LmF1dGgwLmNvbTAeFw0xNzEwMTYwMDM1MjVaFw0z
	MTA2MjUwMDM1MjVaMCMxITAfBgNVBAMTGG9zY2FyaGVhbHRoLmF1LmF1dGgwLmNv
	bTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANW2kVipUpCNpZFlkc8X
	CPocm5R1iNF2efvYHZNmhJDD3FkvHwFU9wQzJo3gUugLHW0k9qsN01I7olPC4lLJ
	rYQHc4FQ9fh6D50WPavsWrwFaxjZJIds/91xiCb5X2LZZtBATdcky5zRNKSN6GWQ
	zyyM46TmCSR9iwfBLjzXvgoDgrVMOcsdArTkpB2mpZ7z6WoAUsJiohd0+U3JsgHd
	apcnuriWZe6E6u7rrxdwaYdgAqaRxIah5946nXlMFIX5hBSq0VbCOe0gVftydMbT
	hVu0qCSlVEE91h4Z4d/bNP/tFLCwVotztJDQhwSrUmVC8Avhjx5p2/Ip2DVscJV8
	jLMCAwEAAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUF2F+CVkOCBE/
	HZnkz9Df46qYlH4wDgYDVR0PAQH/BAQDAgKEMA0GCSqGSIb3DQEBCwUAA4IBAQAN
	0/ZpXQ2c6TatmZaLxlR8Z6GGx6xt6xq7qW9lKXgkcHZZmYfCm5tZ1HQHHwYd5ecC
	XVaNBFHGaIJrH7d5HoWpxI1Onj/VIaDO9PGSEpHswWq8HedWjv8SH4l4zAba0wQE
	3H7rcTgf7sT6GIf+J2Yq7YeRL9pkNAE1k32nSBqlr3niy7Hdf/8BFOadNLeuv7fZ
	Dc62UJrFO0JUfOLd6vf1sdE1GF3JDsfsG90TwzsO4RMXEkqat6TBq6+YDpNYMgts
	qt6mTJTOcee03PHRc0ucQRcU+TQz3GHdsF3T8WO1lsZ+1ysfyzfzr2KAzL+FQsj8
	VUMIDoDRBDVuV3bm+46A
	`
)

func AuthorizeServer() gin.HandlerFunc{
	return func(c *gin.Context){
		//后端client_credentials 模式去请求token
		backToken := requestToken("", "client_credentials")
		fmt.Printf("================成功获取服务器端token:%v\n", backToken.AccessToken)
		// err := checkSignFromJWT(backToken.AccessToken, publicKey)
		// if err != nil{
		// 	fmt.Printf("...................jwt 验证失败")
		// }
		session["back-token"] = backToken.AccessToken
		c.Request.Header.Set("Authorization", "Bearer "+backToken.AccessToken)
		c.Next()
		
	}
}

func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		// 从JWKS中，通过解析请求头中的token(只是解析header和claims部分)，可以得到kid，根据这个kid在jwks中找到对应的jwk，然后就可以得到相应的certificate,即public key 可用于验证签名

		fmt.Println("............Retrieve the JWKS")
		jwksURL := "https://oscarhealth.au.auth0.com/.well-known/jwks.json"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwksURL})
		audience := []string{"http://localhost:9090/brand"}
		configuration := auth0.NewConfiguration(client, audience, "https://oscarhealth.au.auth0.com/", jose.RS256)
		validator := auth0.NewValidator(configuration)
		token, err := validator.ValidateRequest(c.Request)

		if err != nil {
			fmt.Println("Token is not valid or missing token")
			c.JSON(http.StatusUnauthorized, gin.H{"Message":"Missing or invalid token.",})
			c.Abort()
		} else {
			// Ensure the token has the correct scope
			result := checkScope(c.Request, validator, token)
			if result == true {
				// If the token is valid and we have the right scope, we'll pass through the middleware
				c.Next()
			} else {
				fmt.Println("You do not have the read:messages scope.")
				c.JSON(http.StatusUnauthorized, gin.H{"Message":"You do not have the read:messages scope.",})
				c.Abort()
			}
		}
		return
	}
}

func checkScope(r *http.Request, validator *auth0.JWTValidator, token *jwt.JSONWebToken) bool {
	claims := map[string]interface{}{}
	err := validator.Claims(r, token, &claims)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if strings.Contains(claims["scope"].(string), "read:brand") {
		return true
	} else {
		return false
	}
}

//Test method
func (oauthController *OauthController) TestMethod(c *gin.Context) {
	// body := gin.H{
	// 	"grant_type":    "authorization_code",
	// 	"client_id":     conf.Oauth2.ClientID,
	// 	"client_secret": conf.Oauth2.ClientSecret,
	// 	"code":          code,
	// 	"redirect_uri":  conf.Oauth2.RedirectURL,
	// }

	

	// 转发前端传过来的token 
	frontToken := session["front_token"]
	url := "https://oscarhealth.au.auth0.com/userinfo"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+frontToken)


	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("====================userinfo")
	fmt.Println(string(body))
	
	// jwtToken := wrapJWT(backToken.AccessToken)
	// fmt.Println(jwtToken)

	//请求api，set header {"authorization":"Barer token"}
	//callAPI(*token)
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
	//sendJWT(token)

}

func requestToken(code, method string) *Token {

	var token Token

	var resp *http.Response
	switch method {
	case "authorization_code":
		body := "{\"grant_type\":\"authorization_code\",\"client_id\": \"" + conf.Oauth2.ClientID + "\",\"client_secret\": \"" + conf.Oauth2.ClientSecret + "\",\"code\": \"" + code + "\",\"redirect_uri\": \"" + conf.Oauth2.RedirectURL + "\"}"
		payload := strings.NewReader(body)

		req, _ := http.NewRequest("POST", conf.Oauth2.Endpoint.TokenURL, payload)
		req.Header.Add("content-type", "application/json")
		resp, _ = http.DefaultClient.Do(req)
		break
	case "client_credentials":
		payload := strings.NewReader("{\"grant_type\":\"client_credentials\",\"client_id\": \"" + clientID + "\",\"client_secret\": \"" + clientSecret + "\",\"audience\": \"" + audience + "\"}")
		req, _ := http.NewRequest("POST", conf.Oauth2.Endpoint.TokenURL, payload)
		req.Header.Add("content-type", "application/json")
		resp, _ = http.DefaultClient.Do(req)

		break
	}
	defer resp.Body.Close()
	requestBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(requestBody))
	err := json.Unmarshal(requestBody, &token)
	if err != nil {
		fmt.Println(err)
	}
	return &token
}

func requestUserID(token *Token) {

}

func callAPI(token Token) {
	url := "http://localhost:9090/webapi/v1/brand/4"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("content-type", "application/json")
	authorization := "Bearer " + token.TokenID
	req.Header.Add("authorization", authorization)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("-------------------")
	fmt.Println(string(body))

}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenID     string `json:"id_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func wrapJWT(token string) string {
	// mySigningKey := []byte(clientSecret)

	// jwtToken := jwt.New(jwt.SigningMethodHS256)
	// claims := jwtToken.Claims.(jwt.MapClaims)
	// claims["cltTok"] = token
	// claims["name"] = "oscar"

	// tokString, _ := jwtToken.SignedString(mySigningKey)
	// return tokString
	return "ok"
}

// func checkSignFromJWT(tok, key string) error{
// 	// Parse takes the token string and a function for looking up the key. The latter is especially
// 	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
// 	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
// 	// to the callback, providing flexibility.
// 	token, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return []byte(key), nil
// 	})

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		fmt.Println(".............check the signature valid %v",claims["iss"])
// 		return nil
// 	} else {
// 		fmt.Println(err)
// 		return err
// 	}
// }
