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

// OfferController implements the offer resource.
type OfferController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

// NewOfferController creates a offer controller.
func NewOfferController(service *goa.Service) (*OfferController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &OfferController{
		Controller: service.NewController("OfferController"),
		privateKey: privKey,
	}, nil
}

// Create runs the create action.
func (c *OfferController) Create(ctx *app.CreateOfferContext) error {
	payload := ctx.Payload
	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}

	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		item, _ := ItemDB.OneItem(ctx.Context, payload.ItemID, 0, 0, 0)

		if userID, ok := claims["user_id"].(float64); ok {
			offer := models.Offer{}
			offer.UserID = int(userID)
			offer.ItemID = payload.ItemID
			offer.OwnerID = item.UserID
			offer.Price = payload.Price
			offer.StartAt = payload.StartAt
			offer.EndAt = payload.EndAt

			err := OfferDB.Add(ctx.Context, &offer)
			if err != nil {
				return ErrDatabaseError(err)
			}
			return ctx.Created()
		}
	}
	errID := errors.New("id error")
	return ctx.BadRequest(errID)
}

// List runs the list action.
func (c *OfferController) List(ctx *app.ListOfferContext) error {
	payload := ctx.Payload
	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}
	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(float64); ok {
			// input offer id
			var objs []*app.Offer
			if payload.OfferID != 0 {
				offer, _ := OfferDB.OneOfferByID(ctx.Context, payload.OfferID)

				if offer.UserID == int(userID) {
					profile, _ := ProfileDB.OneProfilebyUseID(ctx.Context, offer.OwnerID)
					offer.Profile = profile
					objs = append(objs, offer)
					return ctx.OK(objs)
				}
				if offer.OwnerID == int(userID) {
					profile, _ := ProfileDB.OneProfilebyUseID(ctx.Context, offer.UserID)
					offer.Profile = profile
					objs = append(objs, offer)
					return ctx.OK(objs)
				}
				errID := errors.New("id error")
				return ctx.BadRequest(errID)
			}

			offeres := OfferDB.ListOfferPreload(ctx.Context, int(userID))

			for _, offer := range offeres {
				if offer.UserID == int(userID) {
					profile, _ := ProfileDB.OneProfilebyUseID(ctx.Context, offer.OwnerID)
					offer.Profile = profile
				}
				if offer.OwnerID == int(userID) {
					profile, _ := ProfileDB.OneProfilebyUseID(ctx.Context, offer.UserID)
					offer.Profile = profile
				}
				objs = append(objs, offer)
			}

			return ctx.OK(objs)
		}
	}
	errID := errors.New("id error")
	return ctx.BadRequest(errID)

}

// Accept runs the accept action.
func (c *OfferController) Accept(ctx *app.AcceptOfferContext) error {
	payload := ctx.Payload
	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}

	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(float64); ok {

			offer, err := OfferDB.Get(ctx.Context, payload.OfferID)
			if err == gorm.ErrRecordNotFound {
				return ctx.NotFound()
			}
			if offer.OwnerID != int(userID) {
				errID := errors.New("id error")
				return ctx.BadRequest(errID)
			}

			offer.Accepted = true
			err = OfferDB.Update(ctx, offer)
			if err != nil {
				return ErrDatabaseError(err)
			}
			return ctx.NoContent()
		}
	}
	errID := errors.New("id error")
	return ctx.BadRequest(errID)
}
