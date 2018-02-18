package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("forent", func() {
	Description("This is the forent storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")

		Model("Item", func() {
			RendersTo(ItemMedia)
			Description("Item Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the item Model PK field")
			})
			Field("name", gorma.String, func() {})
			Field("description", gorma.String, func() {})
			Field("price", gorma.Integer, func() {})
			Field("compensation", gorma.Integer, func() {})
			Field("user_id", gorma.Integer, func() {})
			Field("category_id", gorma.Integer, func() {})
			Field("place_id", gorma.Integer, func() {})
			Field("image1", gorma.String, func() {})
			Field("image2", gorma.String, func() {})
			Field("image3", gorma.String, func() {})
			Field("image4", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})
	})
})
