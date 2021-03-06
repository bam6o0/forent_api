// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "forent": article Resource Client
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

// ListArticlePayload is the article list action payload.
type ListArticlePayload struct {
	// article ID
	ArticleID *int `form:"articleID,omitempty" json:"articleID,omitempty" xml:"articleID,omitempty"`
	// category id
	CategoryID *int `form:"categoryID,omitempty" json:"categoryID,omitempty" xml:"categoryID,omitempty"`
	// item ID
	ItemID *int `form:"itemID,omitempty" json:"itemID,omitempty" xml:"itemID,omitempty"`
	// user id
	UserID *int `form:"userID,omitempty" json:"userID,omitempty" xml:"userID,omitempty"`
}

// ListArticlePath computes a request path to the list action of article.
func ListArticlePath() string {

	return fmt.Sprintf("/v1/articles")
}

// Retrieve all items.
func (c *Client) ListArticle(ctx context.Context, path string, payload *ListArticlePayload, contentType string) (*http.Response, error) {
	req, err := c.NewListArticleRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListArticleRequest create the request corresponding to the list action endpoint of the article resource.
func (c *Client) NewListArticleRequest(ctx context.Context, path string, payload *ListArticlePayload, contentType string) (*http.Request, error) {
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
