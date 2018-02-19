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
	middlecategories := MiddleCategoryDB.ListMiddlecategory(ctx.Context, ctx.LargecategoryID)
	return ctx.OK(middlecategories)
}
