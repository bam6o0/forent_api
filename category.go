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

// All runs the all action.
func (c *CategoryController) All(ctx *app.AllCategoryContext) error {
	allcategories := CategoryDB.AllListCategory(ctx.Context)
	return ctx.OK(allcategories)
}

// List runs the list action.
func (c *CategoryController) List(ctx *app.ListCategoryContext) error {
	pay := *ctx.Payload
	categories := CategoryDB.ListCategory(ctx.Context, *pay.MiddlecategoryID)
	return ctx.OK(categories)
}
