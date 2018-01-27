//go:generate goagen bootstrap -d github.com/tutley/cryptopages/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tutley/cryptopages/app"
)

func main() {
	// Create service
	service := goa.New("cryptopages")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

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
