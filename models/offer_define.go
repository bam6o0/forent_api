package models

import (
	"context"
	"forent_api/app"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// ListOfferPreload returns an array of view: default.
func (m *OfferDB) ListOfferPreload(ctx context.Context, itemID int, userID int) []*app.Offer {
	defer goa.MeasureSince([]string{"goa", "db", "offer", "listoffer"}, time.Now())

	var nativeUser []*Offer
	var nativeOwner []*Offer

	var objs []*app.Offer
	errUser := m.Db.Scopes(OfferFilterByItem(itemID, m.Db), OfferFilterByUser(userID, m.Db)).Table(m.TableName()).Preload("Item").Find(&nativeUser).Error

	if errUser != nil {
		goa.LogError(ctx, "error listing Offer", "error", errUser.Error())
		return objs
	}

	for _, t := range nativeOwner {
		objs = append(objs, t.OfferToOffer())
	}

	errOwner := m.Db.Scopes(OfferFilterByItem(itemID, m.Db), OfferFilterByOwner(userID, m.Db)).Table(m.TableName()).Preload("Item").Find(&nativeOwner).Error
	for _, t := range nativeOwner {
		objs = append(objs, t.OfferToOffer())
	}
	if errOwner != nil {
		goa.LogError(ctx, "error listing Offer", "error", errOwner.Error())
		return objs
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
