package main

import (
	"forent_api/app"
	"forent_api/models"

	"github.com/goadesign/goa"
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
	// ItemController_Delete: start_implement

	// Put your logic here

	return nil
	// ItemController_Delete: end_implement
}

// List runs the list action.
func (c *ItemController) List(ctx *app.ListItemContext) error {
	// ItemController_List: start_implement

	// Put your logic here

	res := app.ItemCollection{}
	return ctx.OK(res)
	// ItemController_List: end_implement
}

// Show runs the show action.
func (c *ItemController) Show(ctx *app.ShowItemContext) error {
	// ItemController_Show: start_implement

	// Put your logic here

	res := &app.Item{}
	return ctx.OK(res)
	// ItemController_Show: end_implement
}

// Update runs the update action.
func (c *ItemController) Update(ctx *app.UpdateItemContext) error {
	// ItemController_Update: start_implement

	// Put your logic here

	return nil
	// ItemController_Update: end_implement
}
