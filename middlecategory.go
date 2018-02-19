package main

import (
	"forent_api/app"
	"github.com/goadesign/goa"
)

// MiddlecategoryController implements the middlecategory resource.
type MiddlecategoryController struct {
	*goa.Controller
}

// NewMiddlecategoryController creates a middlecategory controller.
func NewMiddlecategoryController(service *goa.Service) *MiddlecategoryController {
	return &MiddlecategoryController{Controller: service.NewController("MiddlecategoryController")}
}

// List runs the list action.
func (c *MiddlecategoryController) List(ctx *app.ListMiddlecategoryContext) error {
	// MiddlecategoryController_List: start_implement

	// Put your logic here

	res := app.MiddlecategoryCollection{}
	return ctx.OK(res)
	// MiddlecategoryController_List: end_implement
}

// Show runs the show action.
func (c *MiddlecategoryController) Show(ctx *app.ShowMiddlecategoryContext) error {
	// MiddlecategoryController_Show: start_implement

	// Put your logic here

	res := &app.Middlecategory{}
	return ctx.OK(res)
	// MiddlecategoryController_Show: end_implement
}
