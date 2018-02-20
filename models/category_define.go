package models

import (
	"context"
	"forent_api/app"
	"time"

	"github.com/goadesign/goa"
)

// AllListItem returns an array of view: default.
func (m *CategoryDB) AllListCategory(ctx context.Context) []*app.Category {
	defer goa.MeasureSince([]string{"goa", "db", "category", "allcategory"}, time.Now())

	var native []*Category
	var objs []*app.Category
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing All Category", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.CategoryToCategory())
	}

	return objs
}
