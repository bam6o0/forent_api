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
)

// SiginAuthenticationPath computes a request path to the sigin action of authentication.
func SiginAuthenticationPath() string {

	return fmt.Sprintf("/v1/signin")
}

// signup
func (c *Client) SiginAuthentication(ctx context.Context, path string, payload *SignupPayload, contentType string) (*http.Response, error) {
	req, err := c.NewSiginAuthenticationRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSiginAuthenticationRequest create the request corresponding to the sigin action endpoint of the authentication resource.
func (c *Client) NewSiginAuthenticationRequest(ctx context.Context, path string, payload *SignupPayload, contentType string) (*http.Request, error) {
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

// SignupAuthenticationPath computes a request path to the signup action of authentication.
func SignupAuthenticationPath() string {

	return fmt.Sprintf("/v1/signup")
}

// signup and Creates a valid JWT
func (c *Client) SignupAuthentication(ctx context.Context, path string, payload *SignupPayload, contentType string) (*http.Response, error) {
	req, err := c.NewSignupAuthenticationRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSignupAuthenticationRequest create the request corresponding to the signup action endpoint of the authentication resource.
func (c *Client) NewSignupAuthenticationRequest(ctx context.Context, path string, payload *SignupPayload, contentType string) (*http.Request, error) {
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
	if c.APIKeySigner != nil {
		if err := c.APIKeySigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
