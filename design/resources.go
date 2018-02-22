package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

//User
var _ = Resource("user", func() { // Resources group related API endpoints
	BasePath("/users") // together. They map to REST resources for REST
	DefaultMedia(User) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get user by id") // with its path, parameters (both path
		Routing(GET("/:userID"))      // parameters and querystring values) and payload
		Params(func() {               // (shape of the request body).
			Param("userID", Integer, "user ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new user")
		Payload(func() {
			Param("email", String, "email")
			Param("password", String, "password")

			Required("email", "password")
		})
		Response(Created, "/users/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:userID"),
		)
		Description("Change user data")
		Params(func() {
			Param("userID", Integer, "user ID")
		})
		Payload(func() {
			Param("id", Integer, "Unique user ID")
			Param("email", String, "email")
			Param("password", String, "password")

			Required("id")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:userID"),
		)
		Params(func() {
			Param("userID", Integer, "user ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

})

//Profile
var _ = Resource("profile", func() { // Resources group related API endpoints
	BasePath("/profiles") // together. They map to REST resources for REST
	DefaultMedia(Profile) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get profile by id") // with its path, parameters (both path
		Routing(GET("/:profileID"))      // parameters and querystring values) and payload
		Params(func() {                  // (shape of the request body).
			Param("profileID", Integer, "profile ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new profile")
		Payload(func() {
			Param("first_name", String, "first name")
			Param("last_name", String, "last_name")
			Param("user_id", Integer, "user id")
			Param("introduction", String, "user introduciton")
			Param("avatar_image", String, "avatar image url")
			Param("cover_image", String, "cover image url")

			Required("first_name", "last_name", "user_id", "introduction", "avatar_image", "cover_image")
		})
		Response(Created, "/profiles/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:profileID"),
		)
		Description("Change profile data")
		Params(func() {
			Param("profileID", Integer, "profile ID")
		})
		Payload(func() {
			Param("user_id", Integer, "user id")
			Param("first_name", String, "first name")
			Param("last_name", String, "last_name")
			Param("introduction", String, "user introduciton")
			Param("avatar_image", String, "avatar image url")
			Param("cover_image", String, "cover image url")

			Required("user_id")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:profileID"),
		)
		Params(func() {
			Param("profileID", Integer, "profile ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

})

//Verification
var _ = Resource("verification", func() { // Resources group related API endpoints
	BasePath("/verifications") // together. They map to REST resources for REST
	DefaultMedia(Verification) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get verification by id") // with its path, parameters (both path
		Routing(GET("/:verificationID"))      // parameters and querystring values) and payload
		Params(func() {                       // (shape of the request body).
			Param("verificationID", Integer, "verification ID")
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
			Param("user_id", Integer, "user id")
			Param("identification", Boolean, "identification flag")
			Param("email", Boolean, "address flag")
			Param("facebook_id", Integer, "Unique facebook ID")
			Param("google_id", Integer, "Unique google ID")

			Required("user_id", "identification", "email", "facebook_id", "google_id")
		})
		Response(Created, "/verifications/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:verificationID"),
		)
		Description("Change verification data")
		Params(func() {
			Param("verificationID", Integer, "verification ID")
		})
		Payload(func() {
			Param("user_id", Integer, "user id")
			Param("identification", Boolean, "identification flag")
			Param("email", Boolean, "address flag")
			Param("facebook_id", Integer, "Unique facebook ID")
			Param("google_id", Integer, "Unique google ID")

			Required("user_id")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:verificationID"),
		)
		Params(func() {
			Param("verificationID", Integer, "verification ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

})

// Item
var _ = Resource("item", func() { // Resources group related API endpoints
	BasePath("/items") // together. They map to REST resources for REST
	DefaultMedia(Item) // services.

	Action("list", func() {
		Routing(
			GET(""),
		)
		Payload(func() {
			Param("itemID", Integer, "item ID")
			Param("userID", Integer, "user id")
			Param("placeID", Integer, "place id")
			Param("categoryID", Integer, "category id")
		})
		Description("Retrieve all items.")
		Response(OK, CollectionOf(Item))
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new item")
		Payload(func() {
			Param("name", String, "item name")
			Param("description", String, "description of item")
			Param("price", Integer, "price of item")
			Param("compensation", Integer, "compensation of item")
			Param("user_id", Integer, "user ID")
			Param("category_id", Integer, "category ID")
			Param("place_id", Integer, "category ID")
			Param("image1", String, "item image 1")
			Param("image2", String, "item image 2")
			Param("image3", String, "item image 3")
			Param("image4", String, "item image 4")

			Required("name", "description", "price", "compensation", "user_id", "category_id", "place_id", "image1")
		})
		Response(Created, "/items/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT(""),
		)
		Description("Change item data")
		Payload(func() {
			Param("itemID", Integer, "item ID")
			Param("name", String, "Name of item")
			Param("description", String, "description of item")
			Param("price", Integer, "price of item")
			Param("compensation", Integer, "compensation of item")
			Param("user_id", Integer, "user ID")
			Param("category_id", Integer, "Category ID")
			Param("place_id", Integer, "Place ID")
			Param("image1", String, "item image 1")
			Param("image2", String, "item image 2")
			Param("image3", String, "item image 3")
			Param("image4", String, "item image 4")

			Required("itemID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE(""),
		)
		Payload(func() {
			Param("itemID", Integer, "item ID")
			Required("itemID")
		})

		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

})

// Article
var _ = Resource("article", func() { // Resources group related API endpoints
	BasePath("/articles") // together. They map to REST resources for REST
	DefaultMedia(Article) // services.

	Action("list", func() {
		Routing(
			GET(""),
		)
		Payload(func() {
			Param("articleID", Integer, "article ID")
			Param("itemID", Integer, "item ID")
			Param("userID", Integer, "user id")
			Param("categoryID", Integer, "category id")
		})
		Description("Retrieve all items.")
		Response(OK, CollectionOf(Article))
	})

})

//Offer
var _ = Resource("offer", func() { // Resources group related API endpoints
	BasePath("/offers")   // together. They map to REST resources for REST
	DefaultMedia(Profile) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get owner by id") // with its path, parameters (both path
		Routing(GET("/:ownerID"))      // parameters and querystring values) and payload
		Params(func() {                // (shape of the request body).
			Param("ownerID", Integer, "owner ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new offer")
		Payload(func() {
			Param("user_id", Integer, "offer user id")
			Param("item_id", Integer, "item id")
			Param("owner_id", Integer, "item id")
			Param("price", Integer, "offer price")
			Param("start_at", DateTime, "rental start at")
			Param("end_at", DateTime, "rental end at")

			Required("user_id", "item_id", "price", "start_at", "end_at")
		})
		Response(Created, "/offers/[0-9]+")
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
		Description("Retrieve all comments.")
		Response(OK, CollectionOf(Comment))
	})

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get comment by id") // with its path, parameters (both path
		Routing(GET("/:itemID"))         // parameters and querystring values) and payload
		Params(func() {                  // (shape of the request body).
			Param("itemID", Integer, "item ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new comment")
		Payload(func() {
			Param("user_id", Integer, "comment user id")
			Param("item_id", Integer, "item id")
			Param("text", String, "comment text")
			Param("reply_to", Integer, "comment_id or null")

			Required("user_id", "item_id", "text", "reply_to")
		})
		Response(Created, "/items/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

})

// Category
var _ = Resource("category", func() { // Resources group related API endpoints
	BasePath("/categories") // together. They map to REST resources for REST
	DefaultMedia(Category)  // services.

	Action("list", func() {
		Routing(
			GET(""),
		)
		Payload(func() {
			Param("middlecategoryID", Integer, "middlecategory id")
		})
		Description("Retrieve all categories.")
		Response(OK, CollectionOf(Category))
	})

})

// MiddleCategory
var _ = Resource("middlecategory", func() { // Resources group related API endpoints
	BasePath("/middlecategories") // together. They map to REST resources for REST
	DefaultMedia(MiddleCategory)  // services.

	Action("list", func() {
		Routing(
			GET(""),
		)
		Payload(func() {
			Param("largecategoryID", Integer, "largecategory id")
		})
		Description("Retrieve all middlecategories.")
		Response(OK, CollectionOf(MiddleCategory))
	})
})

// LargeCategory
var _ = Resource("largecategory", func() { // Resources group related API endpoints
	BasePath("/largecategories") // together. They map to REST resources for REST
	DefaultMedia(LargeCategory)  // services.

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all largecategories.")
		Response(OK, CollectionOf(LargeCategory))
	})

})
