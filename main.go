//go:generate goagen bootstrap -d github.com/tutley/cryptopages/design

package main

import (
	"context"
	"crypto/rsa"
	"log"
	"net/http"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/tutley/cryptopages/app"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PrivateKey  will be used by the JWT signin handler to create a token
var PrivateKey *rsa.PrivateKey

func main() {
	// Create service
	service := goa.New("cryptopages")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Init the Database
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Fatal("DB Connect error: ", err)
	}
	defer session.Close()
	db := session.DB("cryptopages")
	service.Use(MongoMiddleware(db))

	// Setup JWT Keys and middleware
	validationHandler, _ := goa.NewMiddleware(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		token := jwt.ContextJWT(ctx)
		claims, ok := token.Claims.(jwtgo.MapClaims)
		if !ok {
			return jwt.ErrJWTError("unsupported claims shape")
		}
		if val, ok := claims["scopes"].(string); !ok || val != "api:access" {
			return jwt.ErrJWTError("you do not have api access")
		}
		return nil
	})

	keys, err := LoadJWTPublicKeys()
	if err != nil {
		service.LogError("pubKey", "err", err)
	}
	jwtResolver := jwt.NewSimpleResolver(keys)
	PrivateKey, err = LoadJWTPrivateKey()
	if err != nil {
		service.LogError("privKey", "err", err)
	}
	app.UseJWTMiddleware(service, jwt.New(jwtResolver, validationHandler, app.NewJWTSecurity()))

	// Security middleware used to secure the creation of JWT tokens.
	app.UseSigninBasicAuthMiddleware(service, NewBasicAuthMiddleware())

	// Mount "health" controller
	c := NewHealthController(service)
	app.MountHealthController(service, c)
	// Mount "js" controller
	c2 := NewJsController(service)
	app.MountJsController(service, c2)
	// Mount "jwt" controller
	c3 := NewJWTController(service)
	app.MountJWTController(service, c3)
	// Mount "public" controller
	c4 := NewPublicController(service)
	app.MountPublicController(service, c4)
	// Mount "swagger" controller
	c5 := NewSwaggerController(service)
	app.MountSwaggerController(service, c5)
	// Mount "user" controller
	c6 := NewUserController(service)
	app.MountUserController(service, c6)

	// Start service
	if err := service.ListenAndServe(":8010"); err != nil {
		service.LogError("startup", "err", err)
	}

}

// NewBasicAuthMiddleware creates a middleware that checks for the presence of a basic auth header
// and validates its content.
func NewBasicAuthMiddleware() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log basic auth info
			user, pass, ok := req.BasicAuth()
			var ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
			if !ok {
				goa.LogInfo(ctx, "failed basic auth")
				return ErrUnauthorized("missing auth")
			}
			var lu User
			db := GetDB(ctx)
			err := db.C("users").Find(bson.M{"username": user}).One(&lu)
			if err != nil {
				return ErrUnauthorized("user not found")
			}

			err = lu.CheckPassword(pass)
			if err != nil {
				return ErrUnauthorized("incorrect password")
			}
			newctx := context.WithValue(ctx, UsernameKey, user)
			// Proceed
			return h(newctx, rw, req)
		}
	}
}
