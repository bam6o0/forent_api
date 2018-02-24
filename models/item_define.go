package models

import (
	"context"
	"forent_api/app"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// ListItemOnCity returns an array of view: default.
func (m *ItemDB) ListItemOnCity(ctx context.Context, placeIDs []int64, categoryID int) []*app.Item {
	defer goa.MeasureSince([]string{"goa", "db", "item", "listitem"}, time.Now())

	var native []*Item
	var objs []*app.Item
	err := m.Db.Scopes(ItemFilterByPlacesOnCity(placeIDs, m.Db), ItemFilterByCategory(categoryID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Item", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ItemToItem())
	}

	return objs
}

// ItemFilterByPlacesOnCity is a gorm filter for a Belongs To relationship.
func ItemFilterByPlacesOnCity(placeIDs []int64, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if len(placeIDs) > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("place_id in (?)", placeIDs)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}
