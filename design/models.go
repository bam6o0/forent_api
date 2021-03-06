package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("forent", func() {
	Description("This is the forent storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")

		// User Model
		Model("User", func() {
			RendersTo(User)
			Description("User Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the user Model PK field")
			})
			HasOne("Profile")
			HasOne("Verification")
			Field("email", gorma.String, func() {})
			Field("password", gorma.String, func() {})
			Field("salt", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Profile Model
		Model("Profile", func() {
			RendersTo(Profile)
			Description("Profile Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the profile Model PK field")
			})
			Field("first_name", gorma.String, func() {})
			Field("last_name", gorma.String, func() {})
			Field("introduction", gorma.String, func() {})
			Field("avatar_image", gorma.String, func() {})
			Field("cover_image", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Verification Model
		Model("Verification", func() {
			RendersTo(Verification)
			Description("Verification Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the verification Model PK field")
			})
			Field("identification", gorma.Boolean, func() {})
			Field("email", gorma.Boolean, func() {})
			Field("facebook_id", gorma.Integer, func() {})
			Field("google_id", gorma.Integer, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Item Model
		Model("Item", func() {
			RendersTo(Item)
			Description("Item Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the item Model PK field")
			})
			BelongsTo("User")
			BelongsTo("Category")
			BelongsTo("Place")
			Field("name", gorma.String, func() {})
			Field("description", gorma.String, func() {})
			Field("price", gorma.Integer, func() {})
			Field("compensation", gorma.Integer, func() {})
			Field("image1", gorma.String, func() {})
			Field("image2", gorma.String, func() {})
			Field("image3", gorma.String, func() {})
			Field("image4", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Place Model
		Model("Place", func() {
			RendersTo(Place)
			Description("Place Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the place Model PK field")
			})
			Field("name", gorma.String, func() {})
			Field("latitude", gorma.BigDecimal, func() {})
			Field("longitude", gorma.BigDecimal, func() {})
			Field("google_place_id", gorma.String, func() {})
			Field("city_id", gorma.String, func() {})
			Field("city_name", gorma.String, func() {})
			Field("type", gorma.Integer, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Article Model
		Model("Article", func() {
			RendersTo(Article)
			Description("Article Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the article Model PK field")
			})
			BelongsTo("User")
			BelongsTo("Category")
			BelongsTo("Item")
			Field("title1", gorma.String, func() {})
			Field("title2", gorma.String, func() {})
			Field("title3", gorma.String, func() {})
			Field("title4", gorma.String, func() {})
			Field("title5", gorma.String, func() {})
			Field("body1", gorma.String, func() {})
			Field("body2", gorma.String, func() {})
			Field("body3", gorma.String, func() {})
			Field("body4", gorma.String, func() {})
			Field("body5", gorma.String, func() {})
			Field("introduction", gorma.String, func() {})
			Field("item_description", gorma.String, func() {})
			Field("image1", gorma.String, func() {})
			Field("image2", gorma.String, func() {})
			Field("image3", gorma.String, func() {})
			Field("image4", gorma.String, func() {})
			Field("image5", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Offer Model
		Model("Offer", func() {
			RendersTo(Offer)
			Description("Offer Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the offer Model PK field")
			})
			BelongsTo("User")
			BelongsTo("Item")
			Field("owner_id", gorma.Integer, func() {})
			Field("price", gorma.Integer, func() {})
			Field("start_at", gorma.Timestamp, func() {})
			Field("end_at", gorma.Timestamp, func() {})
			Field("accepted", gorma.Boolean, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
			HasOne("Rental")
		})

		// Message Model
		Model("Message", func() {
			RendersTo(Message)
			Description("Message Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the message Model PK field")
			})
			BelongsTo("User")
			BelongsTo("Offer")
			Field("text", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Rental Model
		Model("Rental", func() {
			RendersTo(Rental)
			Description("Rental Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the rental Model PK field")
			})
			Field("delivered_at", gorma.Timestamp, func() {})
			Field("returned_at", gorma.Timestamp, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Comment Model
		Model("Comment", func() {
			RendersTo(Comment)
			Description("Comment Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the comment Model PK field")
			})
			BelongsTo("User")
			BelongsTo("Item")
			Field("text", gorma.String, func() {})
			Field("reply_to", gorma.Integer, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Review Model
		Model("Review", func() {
			RendersTo(Review)
			Description("Review Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the review Model PK field")
			})
			BelongsTo("User")
			BelongsTo("Item")
			Field("text", gorma.String, func() {})
			Field("score", gorma.Integer, func() {})
			Field("image", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Impression Model
		Model("Impression", func() {
			RendersTo(Impression)
			Description("Impression Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the impression Model PK field")
			})
			BelongsTo("User")
			BelongsTo("Item")
			Field("text", gorma.String, func() {})
			Field("score", gorma.Integer, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// Category Model
		Model("Category", func() {
			RendersTo(Category)
			Description("Category Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the category Model PK field")
			})
			BelongsTo("MiddleCategory")
			Field("name", gorma.String, func() {})
			Field("name_base", gorma.String, func() {})
			Field("name_en", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// MiddleCategory Model
		Model("MiddleCategory", func() {
			RendersTo(MiddleCategory)
			Description("MiddleCategory Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the  middle category Model PK field")
			})
			BelongsTo("LargeCategory")
			Field("name", gorma.String, func() {})
			Field("name_base", gorma.String, func() {})
			Field("name_en", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		// LargeCategory Model
		Model("LargeCategory", func() {
			RendersTo(LargeCategory)
			Description("LargeCategory Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the large category Model PK field")
			})
			Field("name", gorma.String, func() {})
			Field("name_base", gorma.String, func() {})
			Field("name_en", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

	})
})
