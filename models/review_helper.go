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

// ListReview returns an array of view: default.
func (m *ReviewDB) ListReview(ctx context.Context, itemID int, userID int) []*app.Review {
	defer goa.MeasureSince([]string{"goa", "db", "review", "listreview"}, time.Now())

	var native []*Review
	var objs []*app.Review
	err := m.Db.Scopes(ReviewFilterByItem(itemID, m.Db), ReviewFilterByUser(userID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Review", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ReviewToReview())
	}

	return objs
}

// ReviewToReview loads a Review and builds the default view of media type Review.
func (m *Review) ReviewToReview() *app.Review {
	review := &app.Review{}
	review.ID = m.ID
	review.Image = m.Image
	review.ItemID = m.ItemID
	review.Score = m.Score
	review.Text = m.Text
	review.UserID = m.UserID

	return review
}

// OneReview loads a Review and builds the default view of media type Review.
func (m *ReviewDB) OneReview(ctx context.Context, id int, itemID int, userID int) (*app.Review, error) {
	defer goa.MeasureSince([]string{"goa", "db", "review", "onereview"}, time.Now())

	var native Review
	err := m.Db.Scopes(ReviewFilterByItem(itemID, m.Db), ReviewFilterByUser(userID, m.Db)).Table(m.TableName()).Preload("Item").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Review", "error", err.Error())
		return nil, err
	}

	view := *native.ReviewToReview()
	return &view, err
}
