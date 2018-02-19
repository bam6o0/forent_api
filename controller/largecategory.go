package controller

import (
	"forent_api/app"

	"github.com/goadesign/goa"
)

// LargecategoryController implements the largecategory resource.
type LargecategoryController struct {
	*goa.Controller
}

// NewLargecategoryController creates a largecategory controller.
func NewLargecategoryController(service *goa.Service) *LargecategoryController {
	return &LargecategoryController{Controller: service.NewController("LargecategoryController")}
}

// List runs the list action.
func (c *LargecategoryController) List(ctx *app.ListLargecategoryContext) error {
	// LargecategoryController_List: start_implement

	// Put your logic here

	res := app.LargecategoryCollection{}
	return ctx.OK(res)
	// LargecategoryController_List: end_implement
}

// Show runs the show action.
func (c *LargecategoryController) Show(ctx *app.ShowLargecategoryContext) error {
	// LargecategoryController_Show: start_implement

	// Put your logic here

	res := &app.Largecategory{}
	return ctx.OK(res)
	// LargecategoryController_Show: end_implement
}
