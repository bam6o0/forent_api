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
)

// MessageController implements the message resource.
type MessageController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

// NewMessageController creates a message controller.
func NewMessageController(service *goa.Service) (*MessageController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &MessageController{
		Controller: service.NewController("MessageController"),
		privateKey: privKey,
	}, nil
}

// Create runs the create action.
func (c *MessageController) Create(ctx *app.CreateMessageContext) error {
	payload := ctx.Payload
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}

	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		offer, _ := OfferDB.OneOffer(ctx.Context, payload.OfferID, 0, 0)

		if float64(offer.UserID) == claims["user_id"] {
			message := models.Message{}
			message.OfferID = payload.OfferID
			message.UserID = offer.UserID
			message.Text = payload.Text

			err := MessageDB.Add(ctx.Context, &message)
			if err != nil {
				return ErrDatabaseError(err)
			}
			return ctx.Created()
		}
		if float64(offer.OwnerID) == claims["user_id"] {
			message := models.Message{}
			message.OfferID = payload.OfferID
			message.UserID = offer.OwnerID
			message.Text = payload.Text

			err := MessageDB.Add(ctx.Context, &message)
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
func (c *MessageController) List(ctx *app.ListMessageContext) error {
	payload := ctx.Payload
	if payload.OfferID == 0 {
		return ctx.NotFound()
	}
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}
	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		offer, _ := OfferDB.OneOffer(ctx.Context, payload.OfferID, 0, 0)
		if float64(offer.UserID) == claims["user_id"] || float64(offer.OwnerID) == claims["user_id"] {
			messages := MessageDB.ListMessage(ctx.Context, payload.OfferID, 0)
			return ctx.OK(messages)
		}
	}
	errID := errors.New("id error")
	return ctx.BadRequest(errID)
}
