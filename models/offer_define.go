package models

import (
	"context"
	"forent_api/app"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// ListOfferByOwner returns an array of view: default.
func (m *OfferDB) ListOfferByOwner(ctx context.Context, ownerID int) []*app.Offer {
	defer goa.MeasureSince([]string{"goa", "db", "offer", "listoffer"}, time.Now())

	var native []*Offer
	var objs []*app.Offer
	err := m.Db.Scopes(OfferFilterByOwner(ownerID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Offer", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.OfferToOffer())
	}

	return objs
}

// OfferFilterByOwner is a gorm filter for a Belongs To relationship.
func OfferFilterByOwner(ownerID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if ownerID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("owner_id = ?", ownerID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}
