package main

import (
	"forent_api/app"

	"github.com/goadesign/goa"
)

// OfferController implements the offer resource.
type OfferController struct {
	*goa.Controller
}

// NewOfferController creates a offer controller.
func NewOfferController(service *goa.Service) *OfferController {
	return &OfferController{Controller: service.NewController("OfferController")}
}

// Create runs the create action.
func (c *OfferController) Create(ctx *app.CreateOfferContext) error {
	// OfferController_Create: start_implement

	// Put your logic here

	return nil
	// OfferController_Create: end_implement
}

// Show runs the show action.
func (c *OfferController) Show(ctx *app.ShowOfferContext) error {
	// OfferController_Show: start_implement

	// Put your logic here

	res := &app.Profile{}
	return ctx.OK(res)
	// OfferController_Show: end_implement
}
