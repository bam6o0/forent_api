package main

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
	largecategories := LargeCategoryDB.ListLargecategory(ctx.Context)
	return ctx.OK(largecategories)
}
