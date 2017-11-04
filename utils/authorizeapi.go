package utils

import(
	"github.com/gin-gonic/gin"
	// "oscar-health-go/controllers"
	"fmt"
)

const JWKS_URI = "https://oscarhealth.au.auth0.com/.well-known/jwks.json"
const AUTH0_API_ISSUER = "https://oscarhealth.au.auth0.com/"

var AUTH0_API_AUDIENCE = []string{"https://oscarhealth.au.auth0.com/api/v2/"}

func AuthorizeAPIToken(c *gin.Context){
	fmt.Println("....................AuthorizeAPIToken")
	// client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: JWKS_URI})
    // audience := AUTH0_API_AUDIENCE

    // configuration := auth0.NewConfiguration(client, audience, AUTH0_API_ISSUER, jose.RS256)
    // validator := auth0.NewValidator(configuration)

    // token, err := validator.ValidateRequest(r)

    // if err != nil {
    //   fmt.Println("Token is not valid or missing token")

    //   response := Response{
    //     Message: "Missing or invalid token.",
    //   }

    //   w.WriteHeader(http.StatusUnauthorized)
    //   json.NewEncoder(w).Encode(response)

    // } else {
    //   h.ServeHTTP(w, r)
    // }
}
