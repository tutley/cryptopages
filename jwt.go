package main

import (
	"fmt"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	uuid "github.com/satori/go.uuid"
	"github.com/tutley/cryptopages/app"
)

// JWTController implements the jwt resource.
type JWTController struct {
	*goa.Controller
}

// NewJWTController creates a jwt controller.
func NewJWTController(service *goa.Service) *JWTController {
	return &JWTController{Controller: service.NewController("JWTController")}
}

// Signin runs the signin action.
func (c *JWTController) Signin(ctx *app.SigninJWTContext) error {
	// JWTController_Signin: start_implement

	// Put your logic here
	fmt.Println("ENTERED THE JWT LOGIN ACTION")
	username := ctx.Value(UsernameKey).(string)
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10m := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	jti, _ := uuid.NewV4()
	token.Claims = jwtgo.MapClaims{
		"iss":    "cryptopages",     // who creates the token and signs it
		"aud":    "Users",           // to whom the token is intended to be sent
		"exp":    in10m,             // time when the token will expire (10 minutes from now)
		"jti":    jti.String(),      // a unique identifier for the token
		"iat":    time.Now().Unix(), // when the token was issued/created (now)
		"nbf":    2,                 // time before which the token is not yet valid (2 minutes ago)
		"sub":    username,          // the subject/principal is whom the token is about
		"scopes": "api:access",      // token scope - not a standard claim
	}
	signedToken, err := token.SignedString(PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Send response
	return ctx.NoContent()

	// JWTController_Signin: end_implement
}
