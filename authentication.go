package main

import (
	"crypto/rsa"
	"fmt"
	"forent_api/app"
	"forent_api/models"
	"forent_api/utils/crypto"
	"io/ioutil"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

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
	payload := ctx.Payload

	user := models.User{}
	// user
	check := db.Where("email = ?", payload.Email).First(&user)
	exists := check.RecordNotFound()
	if exists == true {
		err := check.Error
		c.Service.LogError("Login User", "err", err)
		return ctx.BadRequest(goa.ErrBadRequest("Invalid email or password"))
	}

	hashedPassword := crypto.HashPassword(payload.Password, user.Salt)
	if user.Password != hashedPassword {
		return ctx.BadRequest(goa.ErrBadRequest("Invalid email or password"))
	}

	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10m := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	uuid, _ := uuid.NewV4()
	token.Claims = jwtgo.MapClaims{
		"iss":     "forent",          // who creates the token and signs it
		"aud":     "forent-user",     // to whom the token is intended to be sent
		"exp":     in10m,             // time when the token will expire (10 minutes from now)
		"jti":     uuid.String(),     // a unique identifier for the token
		"iat":     time.Now().Unix(), // when the token was issued/created (now)
		"nbf":     2,                 // time before which the token is not yet valid (2 minutes ago)
		"scopes":  "api:access",      // token scope - not a standard claim
		"user_id": user.ID,
	}
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Send response
	//return ctx.NoContent()
	res := &app.Success{}
	res.ID = user.ID

	return ctx.OK(res)
}

// Signup runs the signup action.
func (c *AuthenticationController) Signup(ctx *app.SignupAuthenticationContext) error {
	payload := ctx.Payload

	exists, err := CheckEmailExists(db, payload.Email)
	if err != nil {
		c.Service.LogError("Register User", "err", err)
		return ctx.InternalServerError()
	}
	if exists != true {
		return ctx.BadRequest(goa.ErrBadRequest("Email already exists"))
	}

	// generate hashpassword and add userDB
	salt := crypto.GenerateSalt()
	hashedPassword := crypto.HashPassword(payload.Password, salt)
	user := models.User{}
	user.Email = payload.Email
	user.Password = hashedPassword
	user.Salt = salt

	errDb := UserDB.Add(ctx.Context, &user)
	if errDb != nil {
		c.Service.LogError("Register User", "err", errDb)
		return ErrDatabaseError(errDb)
	}
	// Get user id
	db.Where("email = ?", payload.Email).Last(&user)

	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10m := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	uuid, _ := uuid.NewV4()
	token.Claims = jwtgo.MapClaims{
		"iss":     "forent",          // who creates the token and signs it
		"aud":     "forent-user",     // to whom the token is intended to be sent
		"exp":     in10m,             // time when the token will expire (10 minutes from now)
		"jti":     uuid.String(),     // a unique identifier for the token
		"iat":     time.Now().Unix(), // when the token was issued/created (now)
		"nbf":     2,                 // time before which the token is not yet valid (2 minutes ago)
		"scopes":  "api:access",      // token scope - not a standard claim
		"user_id": user.ID,
	}
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Send response
	//return ctx.NoContent()
	res := &app.Success{}
	res.ID = user.ID

	return ctx.OK(res)

}

// CheckEmailExists is check email adress in user db
func CheckEmailExists(db *gorm.DB, email string) (bool, error) {

	user := models.User{}
	c := db.Where("email = ?", email).First(&user)
	exists := c.RecordNotFound()
	if exists != true {
		err := c.Error
		return exists, err
	}
	return exists, nil
}
