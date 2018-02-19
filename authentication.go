package main

import (
	"forent_api/app"

	"github.com/goadesign/goa"
)

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
}

// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service *goa.Service) *AuthenticationController {
	return &AuthenticationController{Controller: service.NewController("AuthenticationController")}
}

// Create runs the create action.
func (c *AuthenticationController) Create(ctx *app.CreateAuthenticationContext) error {
	// AuthenticationController_Create: start_implement

	// Put your logic here

	return nil
	// AuthenticationController_Create: end_implement
}

// Delete runs the delete action.
func (c *AuthenticationController) Delete(ctx *app.DeleteAuthenticationContext) error {
	// AuthenticationController_Delete: start_implement

	// Put your logic here

	return nil
	// AuthenticationController_Delete: end_implement
}

// Show runs the show action.
func (c *AuthenticationController) Show(ctx *app.ShowAuthenticationContext) error {
	// AuthenticationController_Show: start_implement

	// Put your logic here

	res := &app.Authentication{}
	return ctx.OK(res)
	// AuthenticationController_Show: end_implement
}

// Update runs the update action.
func (c *AuthenticationController) Update(ctx *app.UpdateAuthenticationContext) error {
	// AuthenticationController_Update: start_implement

	// Put your logic here

	return nil
	// AuthenticationController_Update: end_implement
}
