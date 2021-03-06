package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// User defines the media type used to render Users.
var User = MediaType("application/vnd.cryptopages.user", func() {
	Description("A cryptopages user")
	Reference(UserPayload)
	Attributes(func() {
		Attribute("href", String, "API href for making requests on the user", func() {
			Example("/users/scubasteve")
		})
		// These are inherited from the Reference base type
		Attribute("username")
		Attribute("name")
		Attribute("password")
		Attribute("email")
		Attribute("location")
		Attribute("coins")
		Attribute("otherCoin")
		Attribute("jobCategory")
		Attribute("skills")
		Attribute("jobDescription")
		Attribute("available")
		Required("username", "href", "name", "jobCategory")
	})
	View("default", func() {
		Attribute("username")
		Attribute("href")
		Attribute("name")
		Attribute("password")
		Attribute("email")
		Attribute("location")
		Attribute("coins")
		Attribute("otherCoin")
		Attribute("jobCategory")
		Attribute("skills")
		Attribute("jobDescription")
		Attribute("available")
	})
	View("tiny", func() {
		Attribute("username")
		Attribute("href")
		Attribute("name")
		Attribute("jobCategory")
	})
})

var Success = MediaType("application/vnd.cryptopages.success", func() {
	Description("The common media type to auth responses")
	TypeName("Success")
	Attributes(func() {
		Attribute("ok", Boolean, "Always true")
		Required("ok")
	})
	View("default", func() {
		Attribute("ok")
	})
})
