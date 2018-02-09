package main

import (
	"errors"
	"fmt"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/tutley/cryptopages/app"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// CheckUsername runs the checkUsername action.
func (c *UserController) CheckUsername(ctx *app.CheckUsernameUserContext) error {
	// UserController_CheckUsername: start_implement

	// Put your logic here
	// grab the db and lookup the user
	db := GetDB(ctx.Context)
	username := ctx.Username
	_, err := FindUserByUsername(username, db)
	if err != nil {
		return ctx.NotFound()
	}

	return ctx.NoContent()

	// UserController_CheckUsername: end_implement
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here
	db := GetDB(ctx.Context)
	// TODO: try slapping payload directly into the db
	p := ctx.Payload

	//verify the username isn't already being used
	user, _ := FindUserByUsername(p.Username, db)
	if user != nil {
		return ctx.BadRequest(errors.New("user: The user already exists"))
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.BadRequest(fmt.Errorf("user: %+v", err))
	}
	p.Password = string(passHash)
	err = db.C("users").Insert(p)
	if err != nil {
		// TODO: send internal server error?
		return ctx.BadRequest(err)
	}
	return ctx.Created()

	// UserController_Create: end_implement
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	// UserController_Delete: start_implement

	// Put your logic here
	// first we're going to check to see that you're the same user that you're trying to delete
	// (Otherwise that's a no-no)
	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}
	claims := token.Claims.(jwtgo.MapClaims)

	// Use the claims to authorize
	subject := claims["sub"].(string)
	fmt.Println("user: ", subject)

	if subject != ctx.Username {
		return ctx.BadRequest(errors.New("You do not own that acount"))
	}
	// grab the db and lookup the user
	db := GetDB(ctx.Context)
	u, err := FindUserByUsername(subject, db)
	if err != nil {
		return ctx.NotFound()
	}
	err = db.C("users").Remove(bson.M{"username": u.Username})
	if err != nil {
		return ctx.BadRequest(err)
	}
	return ctx.NoContent()

	// UserController_Delete: end_implement
}

// Search runs the search action.
func (c *UserController) Search(ctx *app.SearchUserContext) error {
	// UserController_Search: start_implement

	// Put your logic here
	// grab the db
	db := GetDB(ctx.Context)
	// formulate the search query objects

	// do the thing

	users := []*app.CryptopagesUser{}
	err := db.C("users").Find(bson.M{}).All(&users)
	if err != nil {
		return ctx.NotFound()
	}
	return ctx.OK(users)

	// cc := []*app.CryptopagesUser{}
	// for _, user := range *users {
	// 	u, _ := user.MapToCryptopagesUser()
	// 	cc = append(cc, &u)
	// }
	// res := cc
	// return ctx.OK(res)

	// UserController_Search: end_implement
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here
	// grab the db and lookup the user
	db := GetDB(ctx.Context)
	username := ctx.Username
	u, err := FindUserByUsername(username, db)
	if err != nil {
		return ctx.NotFound()
	}
	u.Password = nil
	return ctx.OK(u)

	// UserController_Show: end_implement
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	// Put your logic here
	// first we're going to check to see that you're the same user that you're trying to update
	// (Otherwise that's a no-no)
	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}
	claims := token.Claims.(jwtgo.MapClaims)

	// Use the claims to authorize
	subject := claims["sub"].(string)
	fmt.Println("user: ", subject)

	if subject != ctx.Username {
		return ctx.BadRequest(errors.New("You do not own that acount"))
	}
	// grab the db and lookup the user
	db := GetDB(ctx.Context)
	u, err := FindUserByUsername(subject, db)
	if err != nil {
		return ctx.NotFound()
	}

	// Make the changes
	p := ctx.Payload
	p.Password = u.Password
	err = db.C("users").Update(bson.M{"username": u.Username}, p)
	if err != nil {
		return ctx.BadRequest(err)
	}
	return ctx.NoContent()

	// UserController_Update: end_implement
}
