package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

// User defines the media type used to render users.
var User = MediaType("application/vnd.user+json", func() {
	Description("user")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique user ID")
		Attribute("first_name", String, "first name")
		Attribute("last_name", String, "last_name")
		Attribute("email", String, "email")
		Attribute("password", String, "password")

		Required("id", "first_name", "last_name", "email", "password")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("first_name")
		Attribute("last_name")
		Attribute("email")
		Attribute("password")
	})
})

// Profile defines the media type used to render profoles.
var Profile = MediaType("application/vnd.profile+json", func() {
	Description("profile")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique profile ID")
		Attribute("user_id", Integer, "user id")
		Attribute("introduction", String, "user introduciton")
		Attribute("address", String, "address")
		Attribute("phone", Integer, "phone number")
		Attribute("avatar_image", String, "avatar image url")
		Attribute("cover_image", String, "cover image url")

		Required("id", "user_id", "introduction", "address", "phone", "avatar_image", "cover_image")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("user_id")
		Attribute("introduction")
		Attribute("address")
		Attribute("phone")
		Attribute("avatar_image")
		Attribute("cover_image")
	})
})

// Authentication defines the media type used to render Authentications.
var Authentication = MediaType("application/vnd.authentication+json", func() {
	Description("Authentication")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique auth ID")
		Attribute("user_id", Integer, "user id")
		Attribute("identification", Boolean, "identification flag")
		Attribute("email", Boolean, "address flag")
		Attribute("phone", Boolean, "phone flag")

		Required("id", "user_id", "identification", "email", "phone")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("user_id")
		Attribute("identification")
		Attribute("email")
		Attribute("phone")
	})
})

// Item defines the media type used to render items.
var Item = MediaType("application/vnd.item+json", func() {
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

// Article defines the media type used to render Article.
var Article = MediaType("application/vnd.article+json", func() {
	Description("article")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique article ID")
		Attribute("user_id", Integer, "user ID")
		Attribute("item_id", Integer, "item ID")
		Attribute("category_id", Integer, "category ID")

		Attribute("title1", String, "title 1")
		Attribute("title2", String, "title 2")
		Attribute("title3", String, "title 3")
		Attribute("title4", String, "title 4")
		Attribute("title5", String, "title 5")

		Attribute("body1", String, "body 1")
		Attribute("body2", String, "body 2")
		Attribute("body3", String, "body 3")
		Attribute("body4", String, "body 4")
		Attribute("body5", String, "body 5")

		Attribute("introduction", String, "introduction")

		Attribute("image1", String, "item image 1")
		Attribute("image2", String, "item image 2")
		Attribute("image3", String, "item image 3")
		Attribute("image4", String, "item image 4")
		Attribute("image5", String, "item image 5")

		Attribute("item_description", String, "item description")

		Required(
			"id",
			"user_id",
			"category_id",
			"item_id",
			"title1",
			"title2",
			"title3",
			"title4",
			"title5",
			"body1",
			"body2",
			"body3",
			"body4",
			"body5",
			"introduction",
			"item_description",
			"image1",
			"image2",
			"image3",
			"image4",
			"image5",
		)
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("user_id")
		Attribute("item_id")
		Attribute("category_id")
		Attribute("title1")
		Attribute("title2")
		Attribute("title3")
		Attribute("title4")
		Attribute("title5")
		Attribute("body1")
		Attribute("body2")
		Attribute("body3")
		Attribute("body4")
		Attribute("body5")
		Attribute("introduction")
		Attribute("image1")
		Attribute("image2")
		Attribute("image3")
		Attribute("image4")
		Attribute("image5")
		Attribute("item_description")

	})
})

// Offer defines the media type used to render Offer.
var Offer = MediaType("application/vnd.offer+json", func() {
	Description("Offer")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique offer ID")
		Attribute("user_id", Integer, "offer user id")
		Attribute("item_id", Integer, "item id")
		Attribute("owner_id", Integer, "item id")
		Attribute("price", Integer, "offer price")
		Attribute("start_at", DateTime, "rental start at")
		Attribute("end_at", DateTime, "rental end at")
		Attribute("accepted", Boolean, "offer accept")

		Required("id", "user_id", "item_id", "owner_id", "price", "start_at", "end_at", "accepted")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("user_id")
		Attribute("item_id")
		Attribute("owner_id")
		Attribute("price")
		Attribute("start_at")
		Attribute("end_at")
		Attribute("accepted")
	})
})

// Rental defines the media type used to render Rental.
var Rental = MediaType("application/vnd.rental+json", func() {
	Description("Offer")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique rental ID")
		Attribute("offer_id", Integer, "offer id")
		Attribute("delivered_at", DateTime, "delivered at")
		Attribute("returned_at", DateTime, "returned at")

		Required("id", "offer_id", "delivered_at", "returned_at")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("offer_id")
		Attribute("delivered_at")
		Attribute("returned_at")
	})
})

// Comment defines the media type used to render Comment.
var Comment = MediaType("application/vnd.comment+json", func() {
	Description("Comment")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique comment ID")
		Attribute("user_id", Integer, "comment user id")
		Attribute("item_id", Integer, "item id")
		Attribute("text", String, "comment text")
		Attribute("reply_to", Integer, "comment_id or null")

		Required("id", "user_id", "item_id", "text", "reply_to")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("user_id")
		Attribute("item_id")
		Attribute("text")
		Attribute("reply_to")
	})
})

// Review defines the media type used to render Review.
var Review = MediaType("application/vnd.review+json", func() {
	Description("Review")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique review ID")
		Attribute("user_id", Integer, "review user id")
		Attribute("item_id", Integer, "item id")
		Attribute("text", String, "review text")
		Attribute("score", Integer, "score")
		Attribute("image", String, "image")

		Required("id", "user_id", "item_id", "text", "score", "image")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("user_id")
		Attribute("item_id")
		Attribute("text")
		Attribute("score")
		Attribute("image")
	})
})

// Impression defines the media type used to render Impression.
var Impression = MediaType("application/vnd.impression+json", func() {
	Description("Impression")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique impression ID")
		Attribute("reviewer_id", Integer, "reviewer user id")
		Attribute("reviewee_id", Integer, "reviewee user id")
		Attribute("text", String, "impression text")
		Attribute("score", Integer, "impression score")

		Required("id", "reviewer_id", "reviewee_id", "text", "score")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("reviewer_id")
		Attribute("reviewee_id")
		Attribute("text")
		Attribute("score")
	})
})

// Category defines the media type used to render Category.
var Category = MediaType("application/vnd.category+json", func() {
	Description("Category")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique category ID")
		Attribute("middle_category_id", Integer, "middle_category_id")
		Attribute("name", String, "category name")
		Attribute("name_base", String, "category kana name")
		Attribute("name_en", String, "category english name")

		Required("id", "middle_category_id", "name", "name_base", "name_en")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("middle_category_id")
		Attribute("name")
		Attribute("name_base")
		Attribute("name_en")
	})
})

// MiddleCategory defines the media type used to render MiddleCategory.
var MiddleCategory = MediaType("application/vnd.middlecategory+json", func() {
	Description("Middle Category")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique middle category ID")
		Attribute("large_category_id", Integer, "large category id")
		Attribute("name", String, "category name")
		Attribute("name_base", String, "category kana name")
		Attribute("name_en", String, "category english name")

		Required("id", "large_category_id", "name", "name_base", "name_en")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("large_category_id")
		Attribute("name")
		Attribute("name_base")
		Attribute("name_en")
	})
})

// LargeCategory defines the media type used to render LargeCategory.
var LargeCategory = MediaType("application/vnd.largecategory+json", func() {
	Description("Large Category")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique large category ID")
		Attribute("name", String, "category name")
		Attribute("name_base", String, "category kana name")
		Attribute("name_en", String, "category english name")

		Required("id", "name", "name_base", "name_en")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("name")
		Attribute("name_base")
		Attribute("name_en")
	})
})

// Place defines the media type used to render Place.
var Place = MediaType("application/vnd.place+json", func() {
	Description("Place")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique place ID")
		Attribute("name", String, "place name")
		Attribute("latitude", Number, "place latitude")
		Attribute("longitude", Number, "place longitude")
		Attribute("google_place_id", String, "google place id")
		Attribute("type", Integer, "receive type")

		Required("id", "name", "latitude", "longitude", "google_place_id", "type")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must have a "default" view.
		Attribute("name")
		Attribute("latitude")
		Attribute("longitude")
		Attribute("google_place_id")
		Attribute("type")
	})
})
