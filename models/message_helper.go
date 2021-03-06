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

// ListMessage returns an array of view: default.
func (m *MessageDB) ListMessage(ctx context.Context, offerID int, userID int) []*app.Message {
	defer goa.MeasureSince([]string{"goa", "db", "message", "listmessage"}, time.Now())

	var native []*Message
	var objs []*app.Message
	err := m.Db.Scopes(MessageFilterByOffer(offerID, m.Db), MessageFilterByUser(userID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Message", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.MessageToMessage())
	}

	return objs
}

// MessageToMessage loads a Message and builds the default view of media type Message.
func (m *Message) MessageToMessage() *app.Message {
	message := &app.Message{}
	message.ID = m.ID
	message.OfferID = m.OfferID
	message.Text = m.Text
	message.UserID = m.UserID

	return message
}

// OneMessage loads a Message and builds the default view of media type Message.
func (m *MessageDB) OneMessage(ctx context.Context, id int, offerID int, userID int) (*app.Message, error) {
	defer goa.MeasureSince([]string{"goa", "db", "message", "onemessage"}, time.Now())

	var native Message
	err := m.Db.Scopes(MessageFilterByOffer(offerID, m.Db), MessageFilterByUser(userID, m.Db)).Table(m.TableName()).Preload("Offer").Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Message", "error", err.Error())
		return nil, err
	}

	view := *native.MessageToMessage()
	return &view, err
}
