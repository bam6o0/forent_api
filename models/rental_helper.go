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

// ListRental returns an array of view: default.
func (m *RentalDB) ListRental(ctx context.Context) []*app.Rental {
	defer goa.MeasureSince([]string{"goa", "db", "rental", "listrental"}, time.Now())

	var native []*Rental
	var objs []*app.Rental
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Rental", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.RentalToRental())
	}

	return objs
}

// RentalToRental loads a Rental and builds the default view of media type Rental.
func (m *Rental) RentalToRental() *app.Rental {
	rental := &app.Rental{}
	rental.DeliveredAt = m.DeliveredAt
	rental.ID = m.ID
	rental.OfferID = m.OfferID
	rental.ReturnedAt = m.ReturnedAt

	return rental
}

// OneRental loads a Rental and builds the default view of media type Rental.
func (m *RentalDB) OneRental(ctx context.Context, id int) (*app.Rental, error) {
	defer goa.MeasureSince([]string{"goa", "db", "rental", "onerental"}, time.Now())

	var native Rental
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Rental", "error", err.Error())
		return nil, err
	}

	view := *native.RentalToRental()
	return &view, err
}
