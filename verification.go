package main

import (
	"forent_api/app"

	"github.com/goadesign/goa"
)

// VerificationController implements the verification resource.
type VerificationController struct {
	*goa.Controller
}

// NewVerificationController creates a verification controller.
func NewVerificationController(service *goa.Service) *VerificationController {
	return &VerificationController{Controller: service.NewController("VerificationController")}
}

// Create runs the create action.
func (c *VerificationController) Create(ctx *app.CreateVerificationContext) error {
	// VerificationController_Create: start_implement

	// Put your logic here

	return nil
	// VerificationController_Create: end_implement
}

// Show runs the show action.
func (c *VerificationController) Show(ctx *app.ShowVerificationContext) error {
	// VerificationController_Show: start_implement

	// Put your logic here

	res := &app.Verification{}
	return ctx.OK(res)
	// VerificationController_Show: end_implement
}
