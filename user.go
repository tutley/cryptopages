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

	return nil
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
