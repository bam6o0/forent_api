package main

import (
	"forent_api/app"
	"github.com/goadesign/goa"
)

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
	// ItemController_Create: start_implement

	// Put your logic here

	return nil
	// ItemController_Create: end_implement
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
