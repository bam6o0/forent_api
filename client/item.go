// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "forent": item Resource Client
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

// CreateItemPayload is the item create action payload.
type CreateItemPayload struct {
	// category ID
	CategoryID int `form:"category_id" json:"category_id" xml:"category_id"`
	// compensation of item
	Compensation int `form:"compensation" json:"compensation" xml:"compensation"`
	// description of item
	Description string `form:"description" json:"description" xml:"description"`
	// item image 1
	Image1 string `form:"image1" json:"image1" xml:"image1"`
	// item image 2
	Image2 *string `form:"image2,omitempty" json:"image2,omitempty" xml:"image2,omitempty"`
	// item image 3
	Image3 *string `form:"image3,omitempty" json:"image3,omitempty" xml:"image3,omitempty"`
	// item image 4
	Image4 *string `form:"image4,omitempty" json:"image4,omitempty" xml:"image4,omitempty"`
	// item name
	Name string `form:"name" json:"name" xml:"name"`
	// category ID
	PlaceID int `form:"place_id" json:"place_id" xml:"place_id"`
	// price of item
	Price int `form:"price" json:"price" xml:"price"`
	// user ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// CreateItemPath computes a request path to the create action of item.
func CreateItemPath() string {

	return fmt.Sprintf("/items")
}

// Create new item
func (c *Client) CreateItem(ctx context.Context, path string, payload *CreateItemPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateItemRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateItemRequest create the request corresponding to the create action endpoint of the item resource.
func (c *Client) NewCreateItemRequest(ctx context.Context, path string, payload *CreateItemPayload, contentType string) (*http.Request, error) {
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

// DeleteItemPayload is the item delete action payload.
type DeleteItemPayload struct {
	// item ID
	ItemID int `form:"itemID" json:"itemID" xml:"itemID"`
}

// DeleteItemPath computes a request path to the delete action of item.
func DeleteItemPath() string {

	return fmt.Sprintf("/items")
}

// DeleteItem makes a request to the delete action endpoint of the item resource
func (c *Client) DeleteItem(ctx context.Context, path string, payload *DeleteItemPayload, contentType string) (*http.Response, error) {
	req, err := c.NewDeleteItemRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteItemRequest create the request corresponding to the delete action endpoint of the item resource.
func (c *Client) NewDeleteItemRequest(ctx context.Context, path string, payload *DeleteItemPayload, contentType string) (*http.Request, error) {
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
	req, err := http.NewRequest("DELETE", u.String(), &body)
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

// ListItemPayload is the item list action payload.
type ListItemPayload struct {
	// category id
	CategoryID *int `form:"categoryID,omitempty" json:"categoryID,omitempty" xml:"categoryID,omitempty"`
	// item ID
	ItemID *int `form:"itemID,omitempty" json:"itemID,omitempty" xml:"itemID,omitempty"`
	// place id
	PlaceID *int `form:"placeID,omitempty" json:"placeID,omitempty" xml:"placeID,omitempty"`
	// user id
	UserID *int `form:"userID,omitempty" json:"userID,omitempty" xml:"userID,omitempty"`
}

// ListItemPath computes a request path to the list action of item.
func ListItemPath() string {

	return fmt.Sprintf("/items")
}

// Retrieve all items.
func (c *Client) ListItem(ctx context.Context, path string, payload *ListItemPayload, contentType string) (*http.Response, error) {
	req, err := c.NewListItemRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListItemRequest create the request corresponding to the list action endpoint of the item resource.
func (c *Client) NewListItemRequest(ctx context.Context, path string, payload *ListItemPayload, contentType string) (*http.Request, error) {
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
	return req, nil
}

// UpdateItemPayload is the item update action payload.
type UpdateItemPayload struct {
	// Category ID
	CategoryID *int `form:"category_id,omitempty" json:"category_id,omitempty" xml:"category_id,omitempty"`
	// compensation of item
	Compensation *int `form:"compensation,omitempty" json:"compensation,omitempty" xml:"compensation,omitempty"`
	// description of item
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// item image 1
	Image1 *string `form:"image1,omitempty" json:"image1,omitempty" xml:"image1,omitempty"`
	// item image 2
	Image2 *string `form:"image2,omitempty" json:"image2,omitempty" xml:"image2,omitempty"`
	// item image 3
	Image3 *string `form:"image3,omitempty" json:"image3,omitempty" xml:"image3,omitempty"`
	// item image 4
	Image4 *string `form:"image4,omitempty" json:"image4,omitempty" xml:"image4,omitempty"`
	// item ID
	ItemID int `form:"itemID" json:"itemID" xml:"itemID"`
	// Name of item
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Place ID
	PlaceID *int `form:"place_id,omitempty" json:"place_id,omitempty" xml:"place_id,omitempty"`
	// price of item
	Price *int `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	// user ID
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// UpdateItemPath computes a request path to the update action of item.
func UpdateItemPath() string {

	return fmt.Sprintf("/items")
}

// Change item data
func (c *Client) UpdateItem(ctx context.Context, path string, payload *UpdateItemPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateItemRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateItemRequest create the request corresponding to the update action endpoint of the item resource.
func (c *Client) NewUpdateItemRequest(ctx context.Context, path string, payload *UpdateItemPayload, contentType string) (*http.Request, error) {
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
