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

// Impression Model
type Impression struct {
	ID        int `gorm:"primary_key"` // This is the impression Model PK field
	ItemID    int // Belongs To Item
	Score     int
	Text      string
	UserID    int        // Belongs To User
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
	Item      Item
	User      User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Impression) TableName() string {
	return "impressions"

}

// ImpressionDB is the implementation of the storage interface for
// Impression.
type ImpressionDB struct {
	Db *gorm.DB
}

// NewImpressionDB creates a new storage type.
func NewImpressionDB(db *gorm.DB) *ImpressionDB {
	return &ImpressionDB{Db: db}
}

// DB returns the underlying database.
func (m *ImpressionDB) DB() interface{} {
	return m.Db
}

// ImpressionStorage represents the storage interface.
type ImpressionStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Impression, error)
	Get(ctx context.Context, id int) (*Impression, error)
	Add(ctx context.Context, impression *Impression) error
	Update(ctx context.Context, impression *Impression) error
	Delete(ctx context.Context, id int) error

	ListImpression(ctx context.Context, itemID int, userID int) []*app.Impression
	OneImpression(ctx context.Context, id int, itemID int, userID int) (*app.Impression, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *ImpressionDB) TableName() string {
	return "impressions"

}

// Belongs To Relationships

// ImpressionFilterByItem is a gorm filter for a Belongs To relationship.
func ImpressionFilterByItem(itemID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if itemID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("item_id = ?", itemID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// Belongs To Relationships

// ImpressionFilterByUser is a gorm filter for a Belongs To relationship.
func ImpressionFilterByUser(userID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if userID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Impression as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *ImpressionDB) Get(ctx context.Context, id int) (*Impression, error) {
	defer goa.MeasureSince([]string{"goa", "db", "impression", "get"}, time.Now())

	var native Impression
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Impression
func (m *ImpressionDB) List(ctx context.Context) ([]*Impression, error) {
	defer goa.MeasureSince([]string{"goa", "db", "impression", "list"}, time.Now())

	var objs []*Impression
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *ImpressionDB) Add(ctx context.Context, model *Impression) error {
	defer goa.MeasureSince([]string{"goa", "db", "impression", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Impression", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *ImpressionDB) Update(ctx context.Context, model *Impression) error {
	defer goa.MeasureSince([]string{"goa", "db", "impression", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Impression", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ImpressionDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "impression", "delete"}, time.Now())

	var obj Impression

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Impression", "error", err.Error())
		return err
	}

	return nil
}
