package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

// APIKey defines a security scheme using an API key (shared secret).  The scheme uses the
// "X-Shared-Secret" header to lookup the key.
var APIKey = APIKeySecurity("api_key", func() {
	Header("X-Shared-Secret")
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

// Place
var _ = Resource("place", func() { // Resources group related API endpoints

	BasePath("/places") // together. They map to REST resources for REST
	DefaultMedia(Place) // services.

	Action("list", func() {
		Routing(
			GET(""),
		)
		Payload(func() {
			Param("PlaceID", Integer, "place id")
		})
		Description("Retrieve all places.")
		Response(OK, CollectionOf(Place))
	})

})
