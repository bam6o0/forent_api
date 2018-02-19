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

// Authentication Model
type Authentication struct {
	ID             int `gorm:"primary_key"` // This is the authentication Model PK field
	Email          bool
	Identification bool
	Phone          bool
	UserID         int        // Belongs To User
	CreatedAt      time.Time  // timestamp
	DeletedAt      *time.Time // nullable timestamp (soft delete)
	UpdatedAt      time.Time  // timestamp
	User           User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Authentication) TableName() string {
	return "authentications"

}

// AuthenticationDB is the implementation of the storage interface for
// Authentication.
type AuthenticationDB struct {
	Db *gorm.DB
}

// NewAuthenticationDB creates a new storage type.
func NewAuthenticationDB(db *gorm.DB) *AuthenticationDB {
	return &AuthenticationDB{Db: db}
}

// DB returns the underlying database.
func (m *AuthenticationDB) DB() interface{} {
	return m.Db
}

// AuthenticationStorage represents the storage interface.
type AuthenticationStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Authentication, error)
	Get(ctx context.Context, id int) (*Authentication, error)
	Add(ctx context.Context, authentication *Authentication) error
	Update(ctx context.Context, authentication *Authentication) error
	Delete(ctx context.Context, id int) error

	ListAuthentication(ctx context.Context, userID int) []*app.Authentication
	OneAuthentication(ctx context.Context, id int, userID int) (*app.Authentication, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *AuthenticationDB) TableName() string {
	return "authentications"

}

// Belongs To Relationships

// AuthenticationFilterByUser is a gorm filter for a Belongs To relationship.
func AuthenticationFilterByUser(userID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if userID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Authentication as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *AuthenticationDB) Get(ctx context.Context, id int) (*Authentication, error) {
	defer goa.MeasureSince([]string{"goa", "db", "authentication", "get"}, time.Now())

	var native Authentication
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Authentication
func (m *AuthenticationDB) List(ctx context.Context) ([]*Authentication, error) {
	defer goa.MeasureSince([]string{"goa", "db", "authentication", "list"}, time.Now())

	var objs []*Authentication
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *AuthenticationDB) Add(ctx context.Context, model *Authentication) error {
	defer goa.MeasureSince([]string{"goa", "db", "authentication", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Authentication", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *AuthenticationDB) Update(ctx context.Context, model *Authentication) error {
	defer goa.MeasureSince([]string{"goa", "db", "authentication", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Authentication", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *AuthenticationDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "authentication", "delete"}, time.Now())

	var obj Authentication

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Authentication", "error", err.Error())
		return err
	}

	return nil
}
