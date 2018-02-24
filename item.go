package main

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"forent_api/app"
	"forent_api/models"
	"io/ioutil"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/jinzhu/gorm"
)

// ErrDatabaseError is the error returned when a db query fails.
var ErrDatabaseError = goa.NewErrorClass("db_error", 500)

// ItemController implements the item resource.
type ItemController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

// NewItemController creates a item controller.
func NewItemController(service *goa.Service) (*ItemController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &ItemController{
		Controller: service.NewController("ItemController"),
		privateKey: privKey,
	}, nil
}

// Create runs the create action.
func (c *ItemController) Create(ctx *app.CreateItemContext) error {
	payload := ctx.Payload
	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}

	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		//var authID = float64(payload.UserID)
		if claims["user_id"] != float64(payload.UserID) {
			errID := errors.New("id error")
			return ctx.BadRequest(errID)
		}
	} else {
		errID := errors.New("id error")
		return ctx.BadRequest(errID)
	}

	item := models.Item{}
	item.Name = ctx.Payload.Name
	item.Description = ctx.Payload.Description
	item.Price = ctx.Payload.Price
	item.Compensation = ctx.Payload.Compensation
	item.UserID = ctx.Payload.UserID
	item.CategoryID = ctx.Payload.CategoryID
	item.PlaceID = ctx.Payload.PlaceID
	item.Image1 = ctx.Payload.Image1
	item.Image2 = *ctx.Payload.Image2
	item.Image3 = *ctx.Payload.Image3
	item.Image4 = *ctx.Payload.Image4

	err := ItemDB.Add(ctx.Context, &item)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.Created()
}

// Delete runs the delete action.
func (c *ItemController) Delete(ctx *app.DeleteItemContext) error {
	payload := ctx.Payload
	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}

	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		//var authID = float64(payload.UserID)
		if claims["user_id"] != float64(payload.UserID) {
			errID := errors.New("id error")
			return ctx.BadRequest(errID)
		}
	} else {
		errID := errors.New("id error")
		return ctx.BadRequest(errID)
	}

	pay := *ctx.Payload
	err := ItemDB.Delete(ctx.Context, pay.ItemID)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}

// List runs the list action.
func (c *ItemController) List(ctx *app.ListItemContext) error {
	pay := *ctx.Payload
	if *pay.ItemID != 0 {
		var objs []*app.Item
		item, _ := ItemDB.OneItem(ctx.Context, *pay.ItemID, 0, 0, 0)
		objs = append(objs, item)
		return ctx.OK(objs)
	} else if *ctx.Payload.CityID != "" {
		places := PlaceDB.ListPlaceOnCity(ctx.Context, *pay.CityID)
		placeIDs := []int64{}
		for _, place := range places {
			placeIDs = append(placeIDs, int64(place.ID))
		}
		items := ItemDB.ListItemOnCity(ctx.Context, placeIDs)
		return ctx.OK(items)
	}

	items := ItemDB.ListItem(ctx.Context, *pay.CategoryID, *pay.PlaceID, *pay.UserID)
	return ctx.OK(items)

}

// Update runs the update action.
func (c *ItemController) Update(ctx *app.UpdateItemContext) error {
	payload := ctx.Payload
	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}

	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		//var authID = float64(payload.UserID)
		if claims["user_id"] != float64(payload.UserID) {
			errID := errors.New("id error")
			return ctx.BadRequest(errID)
		}
	} else {
		errID := errors.New("id error")
		return ctx.BadRequest(errID)
	}

	pay := *ctx.Payload
	item, err := ItemDB.Get(ctx.Context, pay.ItemID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}

	item.Name = *ctx.Payload.Name
	item.Description = *ctx.Payload.Description
	item.Price = *ctx.Payload.Price
	item.Compensation = *ctx.Payload.Compensation
	item.UserID = ctx.Payload.UserID
	item.CategoryID = *ctx.Payload.CategoryID
	item.PlaceID = *ctx.Payload.PlaceID
	item.Image1 = *ctx.Payload.Image1
	item.Image2 = *ctx.Payload.Image2
	item.Image3 = *ctx.Payload.Image3
	item.Image4 = *ctx.Payload.Image4
	err = ItemDB.Update(ctx, item)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}
