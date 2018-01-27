package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// JWT defines a security scheme using JWT.  The scheme uses the "Authorization" header to lookup
// the token.  It also defines then scope "api".
var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access") // Define "api:access" scope
})

// BasicAuth defines a security scheme using basic authentication. The scheme protects the "signin"
// action used to create JWTs.
var SigninBasicAuth = BasicAuthSecurity("SigninBasicAuth")

// Resource jwt uses the JWTSecurity security scheme.
var _ = Resource("jwt", func() {
	Description("This resource uses JWT to secure its endpoints")
	DefaultMedia(Success)

	Action("signin", func() {
		Description("Creates a valid JWT")
		Security(SigninBasicAuth)
		Routing(POST("/jwt/signin"))
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
	})

})
