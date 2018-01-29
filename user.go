package main

import (
	"github.com/goadesign/goa"
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

	return nil

	// UserController_Delete: end_implement
}

// Search runs the search action.
func (c *UserController) Search(ctx *app.SearchUserContext) error {
	// UserController_Search: start_implement

	// Put your logic here

	res := app.CryptopagesUserCollection{}
	return ctx.OK(res)

	// UserController_Search: end_implement
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here

	res := &app.CryptopagesUser{}
	return ctx.OK(res)

	// UserController_Show: end_implement
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	// Put your logic here

	return nil

	// UserController_Update: end_implement
}
