package controller

import (
	"forent_api/app"

	"github.com/goadesign/goa"
)

// ArticleController implements the article resource.
type ArticleController struct {
	*goa.Controller
}

// NewArticleController creates a article controller.
func NewArticleController(service *goa.Service) *ArticleController {
	return &ArticleController{Controller: service.NewController("ArticleController")}
}

// List runs the list action.
func (c *ArticleController) List(ctx *app.ListArticleContext) error {
	// ArticleController_List: start_implement

	// Put your logic here

	res := app.ArticleCollection{}
	return ctx.OK(res)
	// ArticleController_List: end_implement
}

// Show runs the show action.
func (c *ArticleController) Show(ctx *app.ShowArticleContext) error {
	// ArticleController_Show: start_implement

	// Put your logic here

	res := &app.Article{}
	return ctx.OK(res)
	// ArticleController_Show: end_implement
}
