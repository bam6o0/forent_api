// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "forent": authentication Resource Client
//
// Command:
// $ goagen
// --design=forent_api/design
// --out=$(GOPATH)/src/forent_api
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateAuthenticationPayload is the authentication create action payload.
type CreateAuthenticationPayload struct {
	// address flag
	Email bool `form:"email" json:"email" xml:"email"`
	// identification flag
	Identification bool `form:"identification" json:"identification" xml:"identification"`
	// phone flag
	Phone bool `form:"phone" json:"phone" xml:"phone"`
	// user id
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// CreateAuthenticationPath computes a request path to the create action of authentication.
func CreateAuthenticationPath() string {

	return fmt.Sprintf("/authentications")
}

// Create new authentication
func (c *Client) CreateAuthentication(ctx context.Context, path string, payload *CreateAuthenticationPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateAuthenticationRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateAuthenticationRequest create the request corresponding to the create action endpoint of the authentication resource.
func (c *Client) NewCreateAuthenticationRequest(ctx context.Context, path string, payload *CreateAuthenticationPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteAuthenticationPath computes a request path to the delete action of authentication.
func DeleteAuthenticationPath(userID int) string {
	param0 := strconv.Itoa(userID)

	return fmt.Sprintf("/authentications/%s", param0)
}

// DeleteAuthentication makes a request to the delete action endpoint of the authentication resource
func (c *Client) DeleteAuthentication(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteAuthenticationRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteAuthenticationRequest create the request corresponding to the delete action endpoint of the authentication resource.
func (c *Client) NewDeleteAuthenticationRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowAuthenticationPath computes a request path to the show action of authentication.
func ShowAuthenticationPath(userID int) string {
	param0 := strconv.Itoa(userID)

	return fmt.Sprintf("/authentications/%s", param0)
}

// Get authentication by id
func (c *Client) ShowAuthentication(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowAuthenticationRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowAuthenticationRequest create the request corresponding to the show action endpoint of the authentication resource.
func (c *Client) NewShowAuthenticationRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateAuthenticationPayload is the authentication update action payload.
type UpdateAuthenticationPayload struct {
	// address flag
	Email *bool `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// identification flag
	Identification *bool `form:"identification,omitempty" json:"identification,omitempty" xml:"identification,omitempty"`
	// phone flag
	Phone *bool `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// user id
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// UpdateAuthenticationPath computes a request path to the update action of authentication.
func UpdateAuthenticationPath(userID int) string {
	param0 := strconv.Itoa(userID)

	return fmt.Sprintf("/authentications/%s", param0)
}

// Change authentication data
func (c *Client) UpdateAuthentication(ctx context.Context, path string, payload *UpdateAuthenticationPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateAuthenticationRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateAuthenticationRequest create the request corresponding to the update action endpoint of the authentication resource.
func (c *Client) NewUpdateAuthenticationRequest(ctx context.Context, path string, payload *UpdateAuthenticationPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
