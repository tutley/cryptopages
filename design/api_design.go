package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cryptopages", func() {
	Title("CryptoPages - Do business with fellow crypto enthusiasts.")
	Description("A directory of crypto enthusiasts offering their services. Written in goa.")
	Scheme("http")
	Host("localhost:8010")
	Contact(func() {
		Name("Tom Utley")
		Email("tom@tomutley.com")
		URL("https://tomutley.com")
	})
	BasePath("/")
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		Headers("Authorization", "Content-Type")
		Expose("Authorization")
		MaxAge(600)
		Credentials()
	})
	Consumes("application/json")
	Produces("application/json")
	Security("jwt")
	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})

})
