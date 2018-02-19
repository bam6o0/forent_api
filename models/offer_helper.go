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

// ListOffer returns an array of view: default.
func (m *OfferDB) ListOffer(ctx context.Context, itemID int, userID int) []*app.Offer {
	defer goa.MeasureSince([]string{"goa", "db", "offer", "listoffer"}, time.Now())

	var native []*Offer
	var objs []*app.Offer
	err := m.Db.Scopes(OfferFilterByItem(itemID, m.Db), OfferFilterByUser(userID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Offer", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.OfferToOffer())
	}

	return objs
}

// OfferToOffer loads a Offer and builds the default view of media type Offer.
func (m *Offer) OfferToOffer() *app.Offer {
	offer := &app.Offer{}
	offer.Accepted = m.Accepted
	offer.EndAt = m.EndAt
	offer.ID = m.ID
	offer.ItemID = m.ItemID
	offer.OwnerID = m.OwnerID
	offer.Price = m.Price
	offer.StartAt = m.StartAt
	offer.UserID = m.UserID

	return offer
}

// OneOffer loads a Offer and builds the default view of media type Offer.
func (m *OfferDB) OneOffer(ctx context.Context, id int, itemID int, userID int) (*app.Offer, error) {
	defer goa.MeasureSince([]string{"goa", "db", "offer", "oneoffer"}, time.Now())

	var native Offer
	err := m.Db.Scopes(OfferFilterByItem(itemID, m.Db), OfferFilterByUser(userID, m.Db)).Table(m.TableName()).Preload("Item").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Offer", "error", err.Error())
		return nil, err
	}

	view := *native.OfferToOffer()
	return &view, err
}
