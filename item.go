package main

import (
	"forent_api/app"
	"forent_api/models"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// ErrDatabaseError is the error returned when a db query fails.
var ErrDatabaseError = goa.NewErrorClass("db_error", 500)

// ItemController implements the item resource.
type ItemController struct {
	*goa.Controller
}

// NewItemController creates a item controller.
func NewItemController(service *goa.Service) *ItemController {
	return &ItemController{Controller: service.NewController("ItemController")}
}

// Create runs the create action.
func (c *ItemController) Create(ctx *app.CreateItemContext) error {
	item := models.Item{}
	item.Name = ctx.Payload.Name
	item.Description = ctx.Payload.Description
	item.Price = ctx.Payload.Price
	item.Compensation = ctx.Payload.Compensation
	item.UserID = ctx.Payload.UserID
	item.CategoryID = ctx.Payload.CategoryID
	item.PlaceID = ctx.Payload.PlaceID
	item.Image1 = ctx.Payload.Image1
	item.Image2 = *ctx.Payload.Image2
	item.Image3 = *ctx.Payload.Image3
	item.Image4 = *ctx.Payload.Image4

	err := ItemDB.Add(ctx.Context, &item)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.Created()
}

// Delete runs the delete action.
func (c *ItemController) Delete(ctx *app.DeleteItemContext) error {
	pay := *ctx.Payload
	err := ItemDB.Delete(ctx.Context, pay.ItemID)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *ItemController) List(ctx *app.ListItemContext) error {
	if *ctx.Payload.ItemID != 0 {
		var objs []*app.Item
		pay := *ctx.Payload
		item, _ := ItemDB.OneItem(ctx.Context, *pay.ItemID, 0, 0, 0)
		objs = append(objs, item)
		return ctx.OK(objs)
	}
	pay := *ctx.Payload
	items := ItemDB.ListItem(ctx.Context, *pay.CategoryID, *pay.PlaceID, *pay.UserID)
	return ctx.OK(items)

}

// Update runs the update action.
func (c *ItemController) Update(ctx *app.UpdateItemContext) error {
	pay := *ctx.Payload
	item, err := ItemDB.Get(ctx.Context, pay.ItemID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}

	item.Name = *ctx.Payload.Name
	item.Description = *ctx.Payload.Description
	item.Price = *ctx.Payload.Price
	item.Compensation = *ctx.Payload.Compensation
	item.UserID = *ctx.Payload.UserID
	item.CategoryID = *ctx.Payload.CategoryID
	item.PlaceID = *ctx.Payload.PlaceID
	item.Image1 = *ctx.Payload.Image1
	item.Image2 = *ctx.Payload.Image2
	item.Image3 = *ctx.Payload.Image3
	item.Image4 = *ctx.Payload.Image4
	err = ItemDB.Update(ctx, item)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}
