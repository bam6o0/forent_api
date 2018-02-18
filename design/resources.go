package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("item", func() { // Resources group related API endpoints
	BasePath("/items")      // together. They map to REST resources for REST
	DefaultMedia(ItemMedia) // services.

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all items.")
		Response(OK, CollectionOf(ItemMedia))
	})

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get item by id") // with its path, parameters (both path
		Routing(GET("/:itemID"))      // parameters and querystring values) and payload
		Params(func() {               // (shape of the request body).
			Param("itemID", Integer, "item ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
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
			PUT("/:itemID"),
		)
		Description("Change item data")
		Params(func() {
			Param("itemID", Integer, "item ID")
		})
		Payload(func() {
			Param("id", Integer, "Unique item ID")
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

			Required("id")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:itemID"),
		)
		Params(func() {
			Param("itemID", Integer, "item ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

})
