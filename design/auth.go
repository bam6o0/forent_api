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
var _ = Resource("signin", func() {
	BasePath("/signin")
	Description("This resource uses JWT to secure its endpoints")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	Action("signin", func() {
		Description("Creates a valid JWT")
		Security(SigninBasicAuth)
		Routing(POST(""))
		Payload(SigninPayload)
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
	})
})

var _ = Resource("signup", func() {
	BasePath("/signup")
	Description("This resource uses JWT to secure its endpoints")

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	Action("signup", func() {
		Description("Creates a valid JWT")
		Security(SigninBasicAuth)
		Routing(POST(""))
		Payload(SignupPayload)
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
	})

})
