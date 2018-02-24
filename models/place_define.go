package models

import (
	"context"
	"forent_api/app"
	"time"

	"github.com/goadesign/goa"
)

// ListPlaceOnCity returns an array of view: default.
func (m *PlaceDB) ListPlaceOnCity(ctx context.Context, cityID string) []*app.Place {
	defer goa.MeasureSince([]string{"goa", "db", "place", "listplace"}, time.Now())

	var native []*Place
	var objs []*app.Place
	err := m.Db.Scopes().Table(m.TableName()).Where("city_id = ?", cityID).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Place", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.PlaceToPlace())
	}

	return objs
}
