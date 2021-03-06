// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "forent": Model Helpers
//
// Command:
// $ goagen
// --design=forent_api/design
// --out=$(GOPATH)/src/forent_api
// --version=v1.3.1

package models

import (
	"context"
	"forent_api/app"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"time"
)

// MediaType Retrieval Functions

// ListUser returns an array of view: default.
func (m *UserDB) ListUser(ctx context.Context) []*app.User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuser"}, time.Now())

	var native []*User
	var objs []*app.User
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUser())
	}

	return objs
}

// UserToUser loads a User and builds the default view of media type User.
func (m *User) UserToUser() *app.User {
	user := &app.User{}
	user.Email = m.Email
	user.ID = m.ID

	return user
}

// OneUser loads a User and builds the default view of media type User.
func (m *UserDB) OneUser(ctx context.Context, id int) (*app.User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuser"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUser()
	return &view, err
}
