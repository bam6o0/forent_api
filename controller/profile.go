package controller

import (
	"forent_api/app"

	"github.com/goadesign/goa"
)

// ProfileController implements the profile resource.
type ProfileController struct {
	*goa.Controller
}

// NewProfileController creates a profile controller.
func NewProfileController(service *goa.Service) *ProfileController {
	return &ProfileController{Controller: service.NewController("ProfileController")}
}

// Create runs the create action.
func (c *ProfileController) Create(ctx *app.CreateProfileContext) error {
	// ProfileController_Create: start_implement

	// Put your logic here

	return nil
	// ProfileController_Create: end_implement
}

// Delete runs the delete action.
func (c *ProfileController) Delete(ctx *app.DeleteProfileContext) error {
	// ProfileController_Delete: start_implement

	// Put your logic here

	return nil
	// ProfileController_Delete: end_implement
}

// Show runs the show action.
func (c *ProfileController) Show(ctx *app.ShowProfileContext) error {
	// ProfileController_Show: start_implement

	// Put your logic here

	res := &app.Profile{}
	return ctx.OK(res)
	// ProfileController_Show: end_implement
}

// Update runs the update action.
func (c *ProfileController) Update(ctx *app.UpdateProfileContext) error {
	// ProfileController_Update: start_implement

	// Put your logic here

	return nil
	// ProfileController_Update: end_implement
}
