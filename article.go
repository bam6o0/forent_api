package main

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
	if *ctx.Payload.ArticleID != 0 {
		var objs []*app.Article
		pay := *ctx.Payload
		article, _ := ArticleDB.OneArticle(ctx.Context, *pay.ArticleID, 0, 0, 0)
		objs = append(objs, article)
		return ctx.OK(objs)
	}
	pay := *ctx.Payload
	articles := ArticleDB.ListArticle(ctx.Context, *pay.CategoryID, *pay.ItemID, *pay.UserID)
	return ctx.OK(articles)
}
