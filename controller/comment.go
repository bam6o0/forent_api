package controller

import (
	"forent_api/app"

	"github.com/goadesign/goa"
)

// CommentController implements the comment resource.
type CommentController struct {
	*goa.Controller
}

// NewCommentController creates a comment controller.
func NewCommentController(service *goa.Service) *CommentController {
	return &CommentController{Controller: service.NewController("CommentController")}
}

// Create runs the create action.
func (c *CommentController) Create(ctx *app.CreateCommentContext) error {
	// CommentController_Create: start_implement

	// Put your logic here

	return nil
	// CommentController_Create: end_implement
}

// List runs the list action.
func (c *CommentController) List(ctx *app.ListCommentContext) error {
	// CommentController_List: start_implement

	// Put your logic here

	res := app.CommentCollection{}
	return ctx.OK(res)
	// CommentController_List: end_implement
}

// Show runs the show action.
func (c *CommentController) Show(ctx *app.ShowCommentContext) error {
	// CommentController_Show: start_implement

	// Put your logic here

	res := &app.Comment{}
	return ctx.OK(res)
	// CommentController_Show: end_implement
}
