package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// UserPayload defines the data structure used in the create user request body.
// It is also the base type for the user media type used to render users.
var UserPayload = Type("UserPayload", func() {
	Attribute("name", String, "The user's full name", func() {
		MinLength(2)
		Example("John Smith")
	})
	Attribute("username", String, "The unique username for this person", func() {
		MinLength(2)
		// TODO: add regex to disallow spaces
		Example("scubasteve")
	})
	Attribute("password", String, "A password (only exposed to user)", func() {
		MinLength(8)
		Example("somethingreallyhardtoguess123%$!@#")
	})
	Attribute("email", func() {
		Description("The user's email address")
		Attribute("value", String, "the email address", func() {
			Format("email")
			MinLength(5)
			Example("me@someplace.com")
		})
		Attribute("makePublic", Boolean, "Should the email address be shown to others")
	})
	Attribute("available", Boolean, "Is this user available to provide work")
	Attribute("location", func() {
		Description("the user's location geographically")
		Attribute("value", String, "the location", func() {
			MinLength(2)
			Example("Charleston, SC, USA")
		})
		Attribute("makePublic", Boolean, "Should the location be shown to others")
	})
	Attribute("coins", func() {
		Description("The coins this user will accept as payment")
		Attribute("btc", Boolean, "Accepts Bitcoin")
		Attribute("ltc", Boolean, "Accepts Litecoin")
		Attribute("bcc", Boolean, "Accepts Bitcoin Cash")
		Attribute("eth", Boolean, "Accepts Ethereum")
		Attribute("neo", Boolean, "Accepts Neo")
		Attribute("xrp", Boolean, "Accepts Ripple")
		Attribute("xlm", Boolean, "Accepts Lumen")
		Attribute("other", func() {
			Description("Accepts some other coin")
			Attribute("name", String, "Name of the other coin accepted")
		})
	})
	Attribute("jobCategory", func() {
		Enum("hardware", "software", "writing", "legal", "labor", "automotive", "services", "others")
	})
	Attribute("skills", ArrayOf(String))
	Attribute("jobDescription", String, func() {
		Example("I'm looking for small remote projects")
	})
})
