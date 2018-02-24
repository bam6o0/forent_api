package models

import (
	"context"
	"forent_api/app"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// OneProfilebyUseID loads a Profile and builds the default view of media type Profile.
func (m *ProfileDB) OneProfilebyUseID(ctx context.Context, id int) (*app.Profile, error) {
	defer goa.MeasureSince([]string{"goa", "db", "profile", "oneprofile"}, time.Now())

	var native Profile
	err := m.Db.Scopes().Table(m.TableName()).Where("user_id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Profile", "error", err.Error())
		return nil, err
	}

	view := *native.ProfileToProfile()
	return &view, err
}
