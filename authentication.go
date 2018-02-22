package main

import (
	"crypto/rsa"
	"fmt"
	"forent_api/app"
	"io/ioutil"

	"github.com/bam6o0/auth_api/repositories"
	"github.com/bam6o0/auth_api/utils/jwt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
)

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

/*
// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service *goa.Service) *AuthenticationController {
	return &AuthenticationController{Controller: service.NewController("AuthenticationController")}
}
*/

// NewAuthenticationController creates a JWT controller.
func NewAuthenticationController(service *goa.Service) (*AuthenticationController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &AuthenticationController{
		Controller: service.NewController("AuthenticationController"),
		privateKey: privKey,
	}, nil
}

// Sigin runs the sigin action.
func (c *AuthenticationController) Sigin(ctx *app.SiginAuthenticationContext) error {
	// AuthenticationController_Sigin: start_implement

	// Put your logic here

	res := &app.Success{}
	return ctx.OK(res)
	// AuthenticationController_Sigin: end_implement
}

// Signup runs the signup action.
func (c *AuthenticationController) Signup(ctx *app.SignupAuthenticationContext) error {
	payload := ctx.Payload
	exists, err := repositories.CheckEmailExists(c.DB, payload.Email)
	if err != nil {
		c.Service.LogError("Register User", "err", err)
		return ctx.InternalServerError()
	}
	if exists {

		return ctx.BadRequest(goa.ErrBadRequest("Email already exists"))
	}
	err = repositories.AddUserToDatabase(c.DB, payload.FirstName, payload.LastName, payload.Email, payload.Password)
	if err != nil {
		c.Service.LogError("Register User", "err", err)
		return ctx.InternalServerError()
	}
	token, err := jwt.CreateJWTToken(payload.Email)
	if err != nil {
		c.Service.LogError("Register User", "err", err)
		return ctx.InternalServerError()
	}
	res := &app.Token{Token: &token}
	return ctx.OK(res)

}
