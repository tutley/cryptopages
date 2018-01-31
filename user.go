package main

import (
	"errors"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/imdario/mergo"
	"github.com/tutley/cryptopages/app"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here
	db := GetDB(ctx.Context)
	p := ctx.Payload

	newEmail := Email{}
	if p.Email != nil {
		newEmail.Value = *p.Email.Value
		newEmail.MakePublic = *p.Email.MakePublic
	}
	newLocation := Location{}
	if p.Location != nil {
		newLocation.Value = *p.Location.Value
		newLocation.MakePublic = *p.Location.MakePublic
	}
	newCoins := Coins{}
	if p.Coins != nil {
		newCoins.BTC = *p.Coins.Btc
		newCoins.LTC = *p.Coins.Ltc
		newCoins.BCC = *p.Coins.Bcc
		newCoins.ETH = *p.Coins.Eth
		newCoins.NEO = *p.Coins.Neo
		newCoins.XRP = *p.Coins.Xrp
		newCoins.XLM = *p.Coins.Xlm
	}
	if p.Coins.Other != nil {
		newCoins.Other = *p.Coins.Other.Name
	}

	newUser := User{
		Name:     p.Name,
		Username: p.Username,
		Password: p.Password,
		Email:    newEmail,
		Location: newLocation,
		Coins:    newCoins,
	}
	if p.Available != nil {
		newUser.Available = *p.Available
	}
	if p.JobCategory != nil {
		newUser.JobCategory = *p.JobCategory
	}
	if len(p.Skills) > 0 {
		newUser.Skills = p.Skills
	}
	if p.JobDescription != nil {
		newUser.JobDescription = *p.JobDescription
	}

	_, err := NewUser(newUser, db)
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
	err = u.Delete(db)
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
	users, err := SearchUsers(db)
	if err != nil {
		return ctx.NotFound()
	}
	cc := []*app.CryptopagesUser{}
	for _, user := range *users {
		u, _ := user.MapToCryptopagesUser()
		cc = append(cc, &u)
	}
	res := cc
	return ctx.OK(res)

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
	cu, err := u.MapToCryptopagesUser()
	if err != nil {
		return ctx.BadRequest(err)
	}
	res := &cu
	return ctx.OK(res)

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
	// TODO: Refactor. This is painful, I must be missing something
	p := ctx.Payload

	newEmail := Email{}
	if p.Email != nil {
		newEmail.Value = *p.Email.Value
		newEmail.MakePublic = *p.Email.MakePublic
	}
	newLocation := Location{}
	if p.Location != nil {
		newLocation.Value = *p.Location.Value
		newLocation.MakePublic = *p.Location.MakePublic
	}
	newCoins := Coins{}
	if p.Coins != nil {
		newCoins.BTC = *p.Coins.Btc
		newCoins.LTC = *p.Coins.Ltc
		newCoins.BCC = *p.Coins.Bcc
		newCoins.ETH = *p.Coins.Eth
		newCoins.NEO = *p.Coins.Neo
		newCoins.XRP = *p.Coins.Xrp
		newCoins.XLM = *p.Coins.Xlm
	}
	if p.Coins.Other != nil {
		newCoins.Other = *p.Coins.Other.Name
	}
	newUser := User{
		Name:     *p.Name,
		Username: *p.Username,
		Password: *p.Password,
		Email:    newEmail,
		Location: newLocation,
		Coins:    newCoins,
	}
	if p.Available != nil {
		newUser.Available = *p.Available
	}
	if p.JobCategory != nil {
		newUser.JobCategory = *p.JobCategory
	}
	if len(p.Skills) > 0 {
		newUser.Skills = p.Skills
	}
	if p.JobDescription != nil {
		newUser.JobDescription = *p.JobDescription
	}
	mergo.Merge(&u, newUser)
	err = u.Save(db)
	if err != nil {
		return ctx.BadRequest(err)
	}
	return ctx.NoContent()

	// UserController_Update: end_implement
}
