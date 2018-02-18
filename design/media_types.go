package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

// ItemMedia defines the media type used to render items.
var ItemMedia = MediaType("application/vnd.item+json", func() {
	Description("A item of user")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique item ID")
		Attribute("name", String, "Name of item")
		Attribute("description", String, "description of item")
		Attribute("price", Integer, "price of item")
		Attribute("compensation", Integer, "compensation of item")
		Attribute("user_id", Integer, "user ID")
		Attribute("category_id", Integer, "Category ID")
		Attribute("place_id", Integer, "Place ID")
		Attribute("image1", String, "item image 1")
		Attribute("image2", String, "item image 2")
		Attribute("image3", String, "item image 3")
		Attribute("image4", String, "item image 4")

		Required("id", "user_id", "name", "description", "price", "compensation", "category_id", "place_id", "image1", "image2", "image3", "image4")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("name")
		Attribute("description")
		Attribute("price")
		Attribute("compensation")
		Attribute("user_id")
		Attribute("category_id")
		Attribute("place_id")
		Attribute("image1")
		Attribute("image2")
		Attribute("image3")
		Attribute("image4")
	})
})
