package main

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"forent_api/app"
	"forent_api/models"
	"io/ioutil"
	"net/http"
	"path/filepath"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
)

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
		//var authID = float64(payload.UserID)
		if claims["user_id"] != float64(payload.UserID) {
			errID := errors.New("id error")
			return ctx.BadRequest(errID)
		}
	} else {
		fmt.Println("失敗")
		errID := errors.New("id error")
		return ctx.BadRequest(errID)
	}

	profile := models.Profile{}
	profile.UserID = payload.UserID
	profile.FirstName = payload.FirstName
	profile.LastName = payload.LastName
	profile.Introduction = ""
	profile.AvatarImage = ""
	profile.CoverImage = ""

	err := ProfileDB.Add(ctx.Context, &profile)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.Created()
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
