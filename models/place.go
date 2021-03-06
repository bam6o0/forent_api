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

// Place Model
type Place struct {
	ID            int `gorm:"primary_key"` // This is the place Model PK field
	CityID        string
	CityName      string
	GooglePlaceID string
	Latitude      float64
	Longitude     float64
	Name          string
	Type          int
	CreatedAt     time.Time  // timestamp
	DeletedAt     *time.Time // nullable timestamp (soft delete)
	UpdatedAt     time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Place) TableName() string {
	return "places"

}

// PlaceDB is the implementation of the storage interface for
// Place.
type PlaceDB struct {
	Db *gorm.DB
}

// NewPlaceDB creates a new storage type.
func NewPlaceDB(db *gorm.DB) *PlaceDB {
	return &PlaceDB{Db: db}
}

// DB returns the underlying database.
func (m *PlaceDB) DB() interface{} {
	return m.Db
}

// PlaceStorage represents the storage interface.
type PlaceStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Place, error)
	Get(ctx context.Context, id int) (*Place, error)
	Add(ctx context.Context, place *Place) error
	Update(ctx context.Context, place *Place) error
	Delete(ctx context.Context, id int) error

	ListPlace(ctx context.Context) []*app.Place
	OnePlace(ctx context.Context, id int) (*app.Place, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *PlaceDB) TableName() string {
	return "places"

}

// CRUD Functions

// Get returns a single Place as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *PlaceDB) Get(ctx context.Context, id int) (*Place, error) {
	defer goa.MeasureSince([]string{"goa", "db", "place", "get"}, time.Now())

	var native Place
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Place
func (m *PlaceDB) List(ctx context.Context) ([]*Place, error) {
	defer goa.MeasureSince([]string{"goa", "db", "place", "list"}, time.Now())

	var objs []*Place
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *PlaceDB) Add(ctx context.Context, model *Place) error {
	defer goa.MeasureSince([]string{"goa", "db", "place", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Place", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *PlaceDB) Update(ctx context.Context, model *Place) error {
	defer goa.MeasureSince([]string{"goa", "db", "place", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Place", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *PlaceDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "place", "delete"}, time.Now())

	var obj Place

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Place", "error", err.Error())
		return err
	}

	return nil
}
