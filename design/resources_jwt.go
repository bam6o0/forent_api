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

	Action("signup", func() {
		Description("signup and Creates a valid JWT")
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

//Profile
var _ = Resource("profile", func() { // Resources group related API endpoints
	BasePath("/profiles")  // together. They map to REST resources for REST
	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})
	DefaultMedia(Profile) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get profile by id") // with its path, parameters (both path
		Routing(GET(""))                 // parameters and querystring values) and payload
		Payload(func() {                 // (shape of the request body).
			Param("user_id", Integer, "user ID")
			Required("user_id")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new profile")
		Payload(func() {
			Param("first_name", String, "first name")
			Param("last_name", String, "last_name")

			Required("first_name", "last_name")
		})
		Response(Created, "/profiles/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

})

//Verification
var _ = Resource("verification", func() { // Resources group related API endpoints
	BasePath("/verifications") // together. They map to REST resources for REST
	DefaultMedia(Verification) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get verification by id") // with its path, parameters (both path
		Routing(GET(""))                      // parameters and querystring values) and payload
		Payload(func() {                      // (shape of the request body).
			Param("user_id", Integer, "user id")
			Required("user_id")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new verification")
		Payload(func() {
			Param("identification", Boolean, "identification flag")
			Param("email", Boolean, "address flag")
			Param("facebook_id", Integer, "Unique facebook ID")
			Param("google_id", Integer, "Unique google ID")

			Required("identification", "email", "facebook_id", "google_id")
		})
		Response(Created, "/verifications/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

})

// Item
var _ = Resource("item", func() { // Resources group related API endpoints

	BasePath("/items") // together. They map to REST resources for REST
	DefaultMedia(Item) // services.
	Action("list", func() {
		Description("Retrieve all items.")
		Routing(
			GET(""),
		)
		Payload(func() {
			Param("itemID", Integer, "item ID")
			Param("userID", Integer, "user id")
			Param("placeID", Integer, "place id")
			Param("categoryID", Integer, "category id")
			Param("cityID", String, "city id")
		})
		Response(OK, CollectionOf(Item))
		Response(NotFound)
	})

	Action("create", func() {
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:access") // Enforce presence of "api" scope in JWT claims.
		})
		Routing(
			POST(""),
		)
		Description("Create new item")
		Payload(func() {
			Param("name", String, "item name")
			Param("description", String, "description of item")
			Param("price", Integer, "price of item")
			Param("compensation", Integer, "compensation of item")
			Param("category_id", Integer, "category ID")
			Param("place_id", Integer, "category ID")
			Param("image1", String, "item image 1")
			Param("image2", String, "item image 2")
			Param("image3", String, "item image 3")
			Param("image4", String, "item image 4")

			Required("name", "description", "price", "compensation", "category_id", "place_id", "image1")
		})
		Response(Created, "/items/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})
})

//Offer
var _ = Resource("offer", func() { // Resources group related API endpoints
	BasePath("/offers")   // together. They map to REST resources for REST
	DefaultMedia(Profile) // services.

	Action("list", func() {
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:access") // Enforce presence of "api" scope in JWT claims.
		})
		Description("Retrieve all offers.")
		Routing(
			GET(""),
		)
		Payload(func() {
			Param("offer_id", Integer, "offer ID")
			Required("offer_id")
		})
		Response(OK, CollectionOf(Offer))
		Response(NotFound) // of HTTP responses.
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:access") // Enforce presence of "api" scope in JWT claims.
		})
		Routing(
			POST(""),
		)
		Description("Create new offer")
		Payload(func() {
			Param("item_id", Integer, "item id")
			Param("price", Integer, "offer price")
			Param("start_at", DateTime, "rental start at")
			Param("end_at", DateTime, "rental end at")

			Required("item_id", "price", "start_at", "end_at")
		})
		Response(Created, "/offers/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("accept", func() {
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:access") // Enforce presence of "api" scope in JWT claims.
		})
		Routing(
			PUT(""),
		)
		Description("accepted offer")
		Payload(func() {
			Param("offer_id", Integer, "offer ID")
			Required("offer_id")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

// Message
var _ = Resource("message", func() { // Resources group related API endpoints
	BasePath("/messages") // together. They map to REST resources for REST
	DefaultMedia(Message) // services.

	Action("list", func() {
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:access") // Enforce presence of "api" scope in JWT claims.
		})
		Routing(
			GET(""),
		)
		Payload(func() {
			Param("offer_id", Integer, "offer id")
			Required("offer_id")
		})
		Description("Retrieve all message in offer.")
		Response(OK, CollectionOf(Message))
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	Action("create", func() {
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:access") // Enforce presence of "api" scope in JWT claims.
		})
		Routing(
			POST(""),
		)
		Description("Create new message")
		Payload(func() {
			Param("offer_id", Integer, "offer id")
			Param("text", String, "comment text")

			Required("offer_id", "text")
		})
		Response(Created, "/messages/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})
})

// Comment
var _ = Resource("comment", func() { // Resources group related API endpoints
	BasePath("/comments") // together. They map to REST resources for REST
	DefaultMedia(Comment) // services.

	Action("list", func() {
		Routing(
			GET(""),
		)
		Payload(func() {
			Param("item_id", Integer, "item id")
			Required("item_id")
		})
		Description("Retrieve all comments.")
		Response(OK, CollectionOf(Comment))
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new comment")
		Payload(func() {
			Param("item_id", Integer, "item id")
			Param("text", String, "comment text")
			Param("reply_to", Integer, "comment_id or null")

			Required("item_id", "text", "reply_to")
		})
		Response(Created, "/items/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

})
