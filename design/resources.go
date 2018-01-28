package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("user", func() {
	BasePath("/user")
	DefaultMedia(User)
	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})
	Action("show", func() {
		Description("Get user by username")
		Routing(GET("/:username"))
		Params(func() {
			Param("username", String, "username", func() {
				MinLength(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("search", func() {
		Description("Search users based on different criteria")
		Routing(GET(""))
		Params(func() {
			Param("skills", String, "Search by Skills")
			Param("coins", String, "Search by crypto coins accepted (abbreviated)")
			Param("jobCategory", String, "Search by job Category")
		})
		Response(OK, func() {
			Media(CollectionOf(User, func() {
				View("default")
				View("tiny")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Register new user")
		NoSecurity()
		Payload(UserPayload, func() {
			Required("name", "username", "password")
		})
		Response(Created, "^/user/[0-9]+$")
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PATCH("/:username"),
		)
		Params(func() {
			Param("username", String)
		})
		Payload(UserPayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:username"),
		)
		Params(func() {
			Param("username", String, "username")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/ui", "public/html/index.html")
})

var _ = Resource("js", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js/*filepath", "public/js")
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "public/swagger/swagger.json")
})
