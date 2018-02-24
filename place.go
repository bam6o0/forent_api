package main

import (
	"forent_api/app"

	"github.com/goadesign/goa"
)

// PlaceController implements the place resource.
type PlaceController struct {
	*goa.Controller
}

// NewPlaceController creates a place controller.
func NewPlaceController(service *goa.Service) *PlaceController {
	return &PlaceController{Controller: service.NewController("PlaceController")}
}

// List runs the list action.
func (c *PlaceController) List(ctx *app.ListPlaceContext) error {
	pay := *ctx.Payload
	if *pay.PlaceID != 0 {
		var objs []*app.Place
		place, _ := PlaceDB.OnePlace(ctx.Context, *pay.PlaceID)
		objs = append(objs, place)
		return ctx.OK(objs)
	}
	places := PlaceDB.ListPlace(ctx.Context)
	return ctx.OK(places)
}
