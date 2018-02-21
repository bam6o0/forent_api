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
	"strconv"
)

// CreateProfilePayload is the profile create action payload.
type CreateProfilePayload struct {
	// avatar image url
	AvatarImage string `form:"avatar_image" json:"avatar_image" xml:"avatar_image"`
	// cover image url
	CoverImage string `form:"cover_image" json:"cover_image" xml:"cover_image"`
	// first name
	FirstName string `form:"first_name" json:"first_name" xml:"first_name"`
	// user introduciton
	Introduction string `form:"introduction" json:"introduction" xml:"introduction"`
	// last_name
	LastName string `form:"last_name" json:"last_name" xml:"last_name"`
	// user id
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// CreateProfilePath computes a request path to the create action of profile.
func CreateProfilePath() string {

	return fmt.Sprintf("/profiles")
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
	return req, nil
}

// DeleteProfilePath computes a request path to the delete action of profile.
func DeleteProfilePath(profileID int) string {
	param0 := strconv.Itoa(profileID)

	return fmt.Sprintf("/profiles/%s", param0)
}

// DeleteProfile makes a request to the delete action endpoint of the profile resource
func (c *Client) DeleteProfile(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteProfileRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteProfileRequest create the request corresponding to the delete action endpoint of the profile resource.
func (c *Client) NewDeleteProfileRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowProfilePath computes a request path to the show action of profile.
func ShowProfilePath(profileID int) string {
	param0 := strconv.Itoa(profileID)

	return fmt.Sprintf("/profiles/%s", param0)
}

// Get profile by id
func (c *Client) ShowProfile(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowProfileRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowProfileRequest create the request corresponding to the show action endpoint of the profile resource.
func (c *Client) NewShowProfileRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateProfilePayload is the profile update action payload.
type UpdateProfilePayload struct {
	// avatar image url
	AvatarImage *string `form:"avatar_image,omitempty" json:"avatar_image,omitempty" xml:"avatar_image,omitempty"`
	// cover image url
	CoverImage *string `form:"cover_image,omitempty" json:"cover_image,omitempty" xml:"cover_image,omitempty"`
	// first name
	FirstName *string `form:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name,omitempty"`
	// user introduciton
	Introduction *string `form:"introduction,omitempty" json:"introduction,omitempty" xml:"introduction,omitempty"`
	// last_name
	LastName *string `form:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty"`
	// user id
	ProfileID int `form:"profile_id" json:"profile_id" xml:"profile_id"`
}

// UpdateProfilePath computes a request path to the update action of profile.
func UpdateProfilePath(profileID int) string {
	param0 := strconv.Itoa(profileID)

	return fmt.Sprintf("/profiles/%s", param0)
}

// Change profile data
func (c *Client) UpdateProfile(ctx context.Context, path string, payload *UpdateProfilePayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateProfileRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateProfileRequest create the request corresponding to the update action endpoint of the profile resource.
func (c *Client) NewUpdateProfileRequest(ctx context.Context, path string, payload *UpdateProfilePayload, contentType string) (*http.Request, error) {
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
