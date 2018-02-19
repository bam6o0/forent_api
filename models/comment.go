// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "forent": Models
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

// Comment Model
type Comment struct {
	ID        int `gorm:"primary_key"` // This is the comment Model PK field
	ItemID    int // has many Comment
	ReplyTo   int
	Text      string
	UserID    int        // has many Comment
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
	Item      Item
	User      User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Comment) TableName() string {
	return "comments"

}

// CommentDB is the implementation of the storage interface for
// Comment.
type CommentDB struct {
	Db *gorm.DB
}

// NewCommentDB creates a new storage type.
func NewCommentDB(db *gorm.DB) *CommentDB {
	return &CommentDB{Db: db}
}

// DB returns the underlying database.
func (m *CommentDB) DB() interface{} {
	return m.Db
}

// CommentStorage represents the storage interface.
type CommentStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Comment, error)
	Get(ctx context.Context, id int) (*Comment, error)
	Add(ctx context.Context, comment *Comment) error
	Update(ctx context.Context, comment *Comment) error
	Delete(ctx context.Context, id int) error

	ListComment(ctx context.Context, itemID int, userID int) []*app.Comment
	OneComment(ctx context.Context, id int, itemID int, userID int) (*app.Comment, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *CommentDB) TableName() string {
	return "comments"

}

// Belongs To Relationships

// CommentFilterByItem is a gorm filter for a Belongs To relationship.
func CommentFilterByItem(itemID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if itemID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("item_id = ?", itemID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// Belongs To Relationships

// CommentFilterByUser is a gorm filter for a Belongs To relationship.
func CommentFilterByUser(userID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if userID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Comment as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *CommentDB) Get(ctx context.Context, id int) (*Comment, error) {
	defer goa.MeasureSince([]string{"goa", "db", "comment", "get"}, time.Now())

	var native Comment
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Comment
func (m *CommentDB) List(ctx context.Context) ([]*Comment, error) {
	defer goa.MeasureSince([]string{"goa", "db", "comment", "list"}, time.Now())

	var objs []*Comment
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *CommentDB) Add(ctx context.Context, model *Comment) error {
	defer goa.MeasureSince([]string{"goa", "db", "comment", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Comment", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *CommentDB) Update(ctx context.Context, model *Comment) error {
	defer goa.MeasureSince([]string{"goa", "db", "comment", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Comment", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *CommentDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "comment", "delete"}, time.Now())

	var obj Comment

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Comment", "error", err.Error())
		return err
	}

	return nil
}
