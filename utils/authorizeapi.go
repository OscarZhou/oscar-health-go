package utils

import(
	"github.com/gin-gonic/gin"
    // "oscar-health-go/controllers"
    "github.com/dgrijalva/jwt-go"
    "fmt"
    "strings"
    "net/http"
)

func AuthorizeAPIToken(c *gin.Context){
    fmt.Println("....................AuthorizeAPIToken")
    
    auth := c.Request.Header.Get("authorization")
    jwtString := strings.Split(auth, " ")[1]
    
    token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
        // Don't forget to validate the alg is what you expect:
        if !(token.Method.Alg() == jwt.SigningMethodRS256.Alg()) {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

    
        // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
        return []byte("RdMzoMjq97r596d3IkanEMD28JCZrWa8OpCH0POf5OB0KH8ArosJx4zzgbYlM0td-92T2qmL0WiiRV4u1r"), nil
    })
    fmt.Println(token)
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        fmt.Println("token is valid")
        fmt.Println(claims)
        c.Next()
    } else {
        fmt.Println("token is invalid")
        fmt.Println(err)
        c.JSON(http.StatusForbidden, nil)
    }
}
