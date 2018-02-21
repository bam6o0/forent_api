//go:generate goagen bootstrap -d forent_api/design

package main

import (
	"forent_api/app"

	"forent_api/models"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var logger log15.Logger

// UserDB is User model
var UserDB *models.UserDB

// ProfileDB is Profile model
var ProfileDB *models.ProfileDB

// VerificationDB is Verification model
var VerificationDB *models.VerificationDB

// ItemDB is Item model
var ItemDB *models.ItemDB

// ArticleDB is Article model
var ArticleDB *models.ArticleDB

// CommentDB is Comment model
var CommentDB *models.CommentDB

// ImpressionDB is Impression model
var ImpressionDB *models.ImpressionDB

// ReviewDB is Review model
var ReviewDB *models.ReviewDB

// OfferDB is Offer model
var OfferDB *models.OfferDB

// PlaceDB is Place model
var PlaceDB *models.PlaceDB

// RentalDB is Rental model
var RentalDB *models.RentalDB

// CategoryDB is Category model
var CategoryDB *models.CategoryDB

// MiddleCategoryDB is MiddleCategory model
var MiddleCategoryDB *models.MiddleCategoryDB

// LargeCategoryDB is LargeCategory model
var LargeCategoryDB *models.LargeCategoryDB

func main() {

	var err error
	db, err = gorm.Open("postgres", "host=localhost user=bam6o0 dbname=forent_api sslmode=disable password=")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	//db.DropTable(&models.Item{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Profile{})
	db.AutoMigrate(&models.Verification{})
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Impression{})
	db.AutoMigrate(&models.Offer{})
	db.AutoMigrate(&models.Rental{})
	db.AutoMigrate(&models.Review{})
	db.AutoMigrate(&models.Place{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.MiddleCategory{})
	db.AutoMigrate(&models.LargeCategory{})

	UserDB = models.NewUserDB(db)
	ProfileDB = models.NewProfileDB(db)
	VerificationDB = models.NewVerificationDB(db)
	ItemDB = models.NewItemDB(db)
	ArticleDB = models.NewArticleDB(db)
	CommentDB = models.NewCommentDB(db)
	ImpressionDB = models.NewImpressionDB(db)
	OfferDB = models.NewOfferDB(db)
	RentalDB = models.NewRentalDB(db)
	ReviewDB = models.NewReviewDB(db)
	PlaceDB = models.NewPlaceDB(db)
	CategoryDB = models.NewCategoryDB(db)
	MiddleCategoryDB = models.NewMiddleCategoryDB(db)
	LargeCategoryDB = models.NewLargeCategoryDB(db)

	db.DB().SetMaxOpenConns(50)

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
	// Mount "category" controller
	c2 := NewCategoryController(service)
	app.MountCategoryController(service, c2)
	// Mount "comment" controller
	c3 := NewCommentController(service)
	app.MountCommentController(service, c3)
	// Mount "item" controller
	c4 := NewItemController(service)
	app.MountItemController(service, c4)
	// Mount "largecategory" controller
	c5 := NewLargecategoryController(service)
	app.MountLargecategoryController(service, c5)
	// Mount "middlecategory" controller
	c6 := NewMiddlecategoryController(service)
	app.MountMiddlecategoryController(service, c6)
	// Mount "offer" controller
	c7 := NewOfferController(service)
	app.MountOfferController(service, c7)
	// Mount "profile" controller
	c8 := NewProfileController(service)
	app.MountProfileController(service, c8)
	// Mount "user" controller
	c9 := NewUserController(service)
	app.MountUserController(service, c9)
	// Mount "verification" controller
	c10 := NewVerificationController(service)
	app.MountVerificationController(service, c10)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
