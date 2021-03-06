// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "forent": profile Resource Client
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
)

// CreateProfilePayload is the profile create action payload.
type CreateProfilePayload struct {
	// first name
	FirstName string `form:"first_name" json:"first_name" xml:"first_name"`
	// last_name
	LastName string `form:"last_name" json:"last_name" xml:"last_name"`
}

// CreateProfilePath computes a request path to the create action of profile.
func CreateProfilePath() string {

	return fmt.Sprintf("/v1/profiles")
}

// Create new profile
func (c *Client) CreateProfile(ctx context.Context, path string, payload *CreateProfilePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateProfileRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateProfileRequest create the request corresponding to the create action endpoint of the profile resource.
func (c *Client) NewCreateProfileRequest(ctx context.Context, path string, payload *CreateProfilePayload, contentType string) (*http.Request, error) {
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
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ShowProfilePayload is the profile show action payload.
type ShowProfilePayload struct {
	// user ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// ShowProfilePath computes a request path to the show action of profile.
func ShowProfilePath() string {

	return fmt.Sprintf("/v1/profiles")
}

// Get profile by id
func (c *Client) ShowProfile(ctx context.Context, path string, payload *ShowProfilePayload, contentType string) (*http.Response, error) {
	req, err := c.NewShowProfileRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowProfileRequest create the request corresponding to the show action endpoint of the profile resource.
func (c *Client) NewShowProfileRequest(ctx context.Context, path string, payload *ShowProfilePayload, contentType string) (*http.Request, error) {
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
	req, err := http.NewRequest("GET", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.APIKeySigner != nil {
		if err := c.APIKeySigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
