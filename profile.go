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

// ProfileController implements the profile resource.
type ProfileController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

// NewProfileController creates a profile controller.
func NewProfileController(service *goa.Service) (*ProfileController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &ProfileController{
		Controller: service.NewController("ProfileController"),
		privateKey: privKey,
	}, nil
}

// Create runs the create action.
func (c *ProfileController) Create(ctx *app.CreateProfileContext) error {
	payload := ctx.Payload
	// Retrieve the token claims
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}

	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {

		// cast
		if userID, ok := claims["user_id"].(float64); ok {
			profile := models.Profile{}
			profile.UserID = int(userID)
			profile.FirstName = payload.FirstName
			profile.LastName = payload.LastName

			err := ProfileDB.Add(ctx.Context, &profile)
			if err != nil {
				return ErrDatabaseError(err)
			}
			return ctx.Created()
		}

	}
	errID := errors.New("id error")
	return ctx.BadRequest(errID)

}

// Show runs the show action.
func (c *ProfileController) Show(ctx *app.ShowProfileContext) error {
	payload := ctx.Payload

	profile, err := ProfileDB.OneProfilebyUseID(ctx.Context, payload.UserID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}

	return ctx.OK(profile)

}
