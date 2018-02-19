package main

import (
	"forent_api/app"
	"github.com/goadesign/goa"
)

// CategoryController implements the category resource.
type CategoryController struct {
	*goa.Controller
}

// NewCategoryController creates a category controller.
func NewCategoryController(service *goa.Service) *CategoryController {
	return &CategoryController{Controller: service.NewController("CategoryController")}
}

// List runs the list action.
func (c *CategoryController) List(ctx *app.ListCategoryContext) error {
	// CategoryController_List: start_implement

	// Put your logic here

	res := app.CategoryCollection{}
	return ctx.OK(res)
	// CategoryController_List: end_implement
}
