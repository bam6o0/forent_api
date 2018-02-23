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

// Resource jwt uses the JWTSecurity security scheme.
var _ = Resource("authentication", func() {
	Description("This resource uses auth to secure its endpoints")
	DefaultMedia(SuccessMedia)

	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	Action("signup", func() {
		Description("signup and Creates a valid JWT")
		NoSecurity()
		Routing(POST("/signup"))
		Payload(SignupPayload)
		Response(OK, SuccessMedia, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})

	Action("sigin", func() {
		Description("signup")
		NoSecurity()
		Routing(GET("/signin"))
		Payload(SignupPayload)
		Response(OK, SuccessMedia, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})

})

// SuccessMedia use auth
var SuccessMedia = MediaType("application/vnd.security.success", func() {
	Description("The common media type to all request responses for this example")
	TypeName("Success")
	Attributes(func() {
		Attribute("id", Integer, "Unique user ID")
		Required("id")
	})
	View("default", func() {
		Attribute("id")
	})
})
