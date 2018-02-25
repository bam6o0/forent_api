//go:generate goagen bootstrap -d forent_api/design

package main

import (
	"context"
	"fmt"
	"forent_api/app"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"forent_api/models"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/jwt"
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

// MessageDB is Message model
var MessageDB *models.MessageDB

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
)

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
	db.AutoMigrate(&models.Message{})

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
	MessageDB = models.NewMessageDB(db)

	db.DB().SetMaxOpenConns(50)

	// Create service
	service := goa.New("forent")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security middlewares
	jwtMiddleware, err := NewJWTMiddleware()
	exitOnFailure(err)
	app.UseJWTMiddleware(service, jwtMiddleware)
	app.UseAPIKeyMiddleware(service, NewAPIKeyMiddleware())

	// Mount "article" controller
	c := NewArticleController(service)
	app.MountArticleController(service, c)
	// Mount "authentication" controller
	c2, err := NewAuthenticationController(service)
	exitOnFailure(err)
	app.MountAuthenticationController(service, c2)
	// Mount "category" controller
	c3 := NewCategoryController(service)
	app.MountCategoryController(service, c3)
	// Mount "comment" controller
	c4 := NewCommentController(service)
	app.MountCommentController(service, c4)
	// Mount "item" controller
	c5, err := NewItemController(service)
	exitOnFailure(err)
	app.MountItemController(service, c5)
	// Mount "largecategory" controller
	c6 := NewLargecategoryController(service)
	app.MountLargecategoryController(service, c6)
	// Mount "middlecategory" controller
	c7 := NewMiddlecategoryController(service)
	app.MountMiddlecategoryController(service, c7)
	// Mount "offer" controller
	c8, err := NewOfferController(service)
	exitOnFailure(err)
	app.MountOfferController(service, c8)
	// Mount "profile" controller
	c9, err := NewProfileController(service)
	exitOnFailure(err)
	app.MountProfileController(service, c9)
	// Mount "verification" controller
	c10 := NewVerificationController(service)
	app.MountVerificationController(service, c10)
	// Mount "place" controller
	c11 := NewPlaceController(service)
	app.MountPlaceController(service, c11)
	// Mount "message" controller
	c12, err := NewMessageController(service)
	exitOnFailure(err)
	app.MountMessageController(service, c12)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

// exitOnFailure prints a fatal error message and exits the process with status 1.
func exitOnFailure(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "[CRIT] %s", err.Error())
	os.Exit(1)
}

// NewJWTMiddleware is JWT auth
func NewJWTMiddleware() (goa.Middleware, error) {

	validationHandler, _ := goa.NewMiddleware(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		token := jwt.ContextJWT(ctx)
		claims, ok := token.Claims.(jwtgo.MapClaims)
		if !ok {
			return jwt.ErrJWTError("unsupported claims shape")
		}
		if val, ok := claims["iss"].(string); !ok || val != "forent" {
			return jwt.ErrJWTError("this is not forent api!")
		}
		return nil
	})
	keys, err := LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	return jwt.New(jwt.NewSimpleResolver(keys), validationHandler, app.NewJWTSecurity()), nil
}

// NewAPIKeyMiddleware creates a middleware that checks for the presence of an authorization header
// and validates its content.
func NewAPIKeyMiddleware() goa.Middleware {
	// Instantiate API Key security scheme details generated from design
	scheme := app.NewAPIKeySecurity()

	apikey, _ := ioutil.ReadFile("./jwtkey/api.key")
	fmt.Println(apikey)

	// Middleware
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log header specified by scheme

			key := req.Header.Get(scheme.Name)
			if string(apikey) != key {
				goa.LogInfo(ctx, "failed api key auth")
				return ErrUnauthorized("missing auth")
			}
			// A real app would do something more interesting here
			if len(key) == 0 || key == "Bearer" {
				goa.LogInfo(ctx, "failed api key auth")
				return ErrUnauthorized("missing auth")
			}
			// Proceed.
			goa.LogInfo(ctx, "auth", "apikey", "key", key)
			return h(ctx, rw, req)
		}
	}
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func LoadJWTPublicKeys() ([]jwt.Key, error) {
	keyFiles, err := filepath.Glob("./jwtkey/*.pub")
	if err != nil {
		return nil, err
	}
	keys := make([]jwt.Key, len(keyFiles))
	for i, keyFile := range keyFiles {
		pem, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return nil, err
		}
		key, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
		}
		keys[i] = key
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("couldn't load public keys for JWT security")
	}

	return keys, nil
}
