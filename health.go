package main

import (
	"github.com/goadesign/goa"
	"github.com/tutley/cryptopages/app"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service) *HealthController {
	return &HealthController{Controller: service.NewController("HealthController")}
}

// Health runs the health action.
func (c *HealthController) Health(ctx *app.HealthHealthContext) error {
	// HealthController_Health: start_implement

	// Put your logic here
	// if _, ok := b.db.GetAccount(1); !ok {
	// 	return fmt.Errorf("failed to connect to DB")
	// }
	// return c.OK([]byte("ok"))
	return nil

	// HealthController_Health: end_implement
}
