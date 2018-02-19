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

// Profile Model
type Profile struct {
	ID           int `gorm:"primary_key"` // This is the profile Model PK field
	Address      string
	AvatarImage  string
	CoverImage   string
	Introduction string
	Phone        string
	UserID       int        // Belongs To User
	CreatedAt    time.Time  // timestamp
	DeletedAt    *time.Time // nullable timestamp (soft delete)
	UpdatedAt    time.Time  // timestamp
	User         User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Profile) TableName() string {
	return "profiles"

}

// ProfileDB is the implementation of the storage interface for
// Profile.
type ProfileDB struct {
	Db *gorm.DB
}

// NewProfileDB creates a new storage type.
func NewProfileDB(db *gorm.DB) *ProfileDB {
	return &ProfileDB{Db: db}
}

// DB returns the underlying database.
func (m *ProfileDB) DB() interface{} {
	return m.Db
}

// ProfileStorage represents the storage interface.
type ProfileStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Profile, error)
	Get(ctx context.Context, id int) (*Profile, error)
	Add(ctx context.Context, profile *Profile) error
	Update(ctx context.Context, profile *Profile) error
	Delete(ctx context.Context, id int) error

	ListProfile(ctx context.Context, userID int) []*app.Profile
	OneProfile(ctx context.Context, id int, userID int) (*app.Profile, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *ProfileDB) TableName() string {
	return "profiles"

}

// Belongs To Relationships

// ProfileFilterByUser is a gorm filter for a Belongs To relationship.
func ProfileFilterByUser(userID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if userID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Profile as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *ProfileDB) Get(ctx context.Context, id int) (*Profile, error) {
	defer goa.MeasureSince([]string{"goa", "db", "profile", "get"}, time.Now())

	var native Profile
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Profile
func (m *ProfileDB) List(ctx context.Context) ([]*Profile, error) {
	defer goa.MeasureSince([]string{"goa", "db", "profile", "list"}, time.Now())

	var objs []*Profile
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *ProfileDB) Add(ctx context.Context, model *Profile) error {
	defer goa.MeasureSince([]string{"goa", "db", "profile", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Profile", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *ProfileDB) Update(ctx context.Context, model *Profile) error {
	defer goa.MeasureSince([]string{"goa", "db", "profile", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Profile", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ProfileDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "profile", "delete"}, time.Now())

	var obj Profile

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Profile", "error", err.Error())
		return err
	}

	return nil
}
