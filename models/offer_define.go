package models

import (
	"context"
	"forent_api/app"
	"time"

	"github.com/goadesign/goa"
)

// ListOfferPreload returns an array of view: default.
func (m *OfferDB) ListOfferPreload(ctx context.Context, itemID int, userID int) []*app.Offer {
	defer goa.MeasureSince([]string{"goa", "db", "offer", "listoffer"}, time.Now())

	var native []*Offer
	var objs []*app.Offer
	err := m.Db.Scopes(OfferFilterByItem(itemID, m.Db), OfferFilterByUser(userID, m.Db)).Table(m.TableName()).Preload("Item").Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Offer", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.OfferToOffer())
	}

	return objs
}
