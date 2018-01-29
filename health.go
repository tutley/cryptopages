package main

import (
	"fmt"
	"github.com/goadesign/goa"
	"github.com/tutley/cryptopages/app"
	"gopkg.in/mgo.v2/bson"
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
	var test map[string]interface{}
	db := GetDB(ctx)
	err := db.C("users").Find(bson.M{}).One(&test)
	if err != nil {
		return fmt.Errorf("failed to connect to DB")
	}
	return ctx.OK([]byte("ok"))

	// HealthController_Health: end_implement
}
