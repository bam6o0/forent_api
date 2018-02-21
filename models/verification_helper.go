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

// ListVerification returns an array of view: default.
func (m *VerificationDB) ListVerification(ctx context.Context) []*app.Verification {
	defer goa.MeasureSince([]string{"goa", "db", "verification", "listverification"}, time.Now())

	var native []*Verification
	var objs []*app.Verification
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Verification", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.VerificationToVerification())
	}

	return objs
}

// VerificationToVerification loads a Verification and builds the default view of media type Verification.
func (m *Verification) VerificationToVerification() *app.Verification {
	verification := &app.Verification{}
	verification.Email = m.Email
	verification.FacebookID = m.FacebookID
	verification.GoogleID = m.GoogleID
	verification.ID = m.ID
	verification.Identification = m.Identification

	return verification
}

// OneVerification loads a Verification and builds the default view of media type Verification.
func (m *VerificationDB) OneVerification(ctx context.Context, id int) (*app.Verification, error) {
	defer goa.MeasureSince([]string{"goa", "db", "verification", "oneverification"}, time.Now())

	var native Verification
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Verification", "error", err.Error())
		return nil, err
	}

	view := *native.VerificationToVerification()
	return &view, err
}
