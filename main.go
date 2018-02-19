//go:generate goagen bootstrap -d forent_api/design

package main

import (
	"forent_api/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("forent")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "article" controller
	c := NewArticleController(service)
	app.MountArticleController(service, c)
	// Mount "authentication" controller
	c2 := NewAuthenticationController(service)
	app.MountAuthenticationController(service, c2)
	// Mount "category" controller
	c3 := NewCategoryController(service)
	app.MountCategoryController(service, c3)
	// Mount "comment" controller
	c4 := NewCommentController(service)
	app.MountCommentController(service, c4)
	// Mount "item" controller
	c5 := NewItemController(service)
	app.MountItemController(service, c5)
	// Mount "largecategory" controller
	c6 := NewLargecategoryController(service)
	app.MountLargecategoryController(service, c6)
	// Mount "middlecategory" controller
	c7 := NewMiddlecategoryController(service)
	app.MountMiddlecategoryController(service, c7)
	// Mount "offer" controller
	c8 := NewOfferController(service)
	app.MountOfferController(service, c8)
	// Mount "profile" controller
	c9 := NewProfileController(service)
	app.MountProfileController(service, c9)
	// Mount "user" controller
	c10 := NewUserController(service)
	app.MountUserController(service, c10)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
