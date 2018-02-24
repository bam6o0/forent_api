// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "forent": place Resource Client
//
// Command:
// $ goagen
// --design=forent_api/design
// --out=$(GOPATH)/src/forent_api
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListPlacePath computes a request path to the list action of place.
func ListPlacePath() string {

	return fmt.Sprintf("/places")
}

// Retrieve all places.
func (c *Client) ListPlace(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListPlaceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListPlaceRequest create the request corresponding to the list action endpoint of the place resource.
func (c *Client) NewListPlaceRequest(ctx context.Context, path string) (*http.Request, error) {
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
