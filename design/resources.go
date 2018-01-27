package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("user", func() {
	BasePath("/user")
	DefaultMedia(User)

	Action("show", func() {
		Description("Get user by username")
		Routing(GET("/:username"))
		Params(func() {
			Param("username", String, "username")
		})
		Response(OK)
		Response(NotFound)
	})
})
