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

// Rental Model
type Rental struct {
	ID          int        `gorm:"primary_key"` // This is the rental Model PK field
	OfferID     int        // has one Rental
	CreatedAt   time.Time  // timestamp
	DeletedAt   *time.Time // nullable timestamp (soft delete)
	DeliveredAt time.Time  // timestamp
	ReturnedAt  time.Time  // timestamp
	UpdatedAt   time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Rental) TableName() string {
	return "rentals"

}

// RentalDB is the implementation of the storage interface for
// Rental.
type RentalDB struct {
	Db *gorm.DB
}

// NewRentalDB creates a new storage type.
func NewRentalDB(db *gorm.DB) *RentalDB {
	return &RentalDB{Db: db}
}

// DB returns the underlying database.
func (m *RentalDB) DB() interface{} {
	return m.Db
}

// RentalStorage represents the storage interface.
type RentalStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Rental, error)
	Get(ctx context.Context, id int) (*Rental, error)
	Add(ctx context.Context, rental *Rental) error
	Update(ctx context.Context, rental *Rental) error
	Delete(ctx context.Context, id int) error

	ListRental(ctx context.Context) []*app.Rental
	OneRental(ctx context.Context, id int) (*app.Rental, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *RentalDB) TableName() string {
	return "rentals"

}

// CRUD Functions

// Get returns a single Rental as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *RentalDB) Get(ctx context.Context, id int) (*Rental, error) {
	defer goa.MeasureSince([]string{"goa", "db", "rental", "get"}, time.Now())

	var native Rental
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Rental
func (m *RentalDB) List(ctx context.Context) ([]*Rental, error) {
	defer goa.MeasureSince([]string{"goa", "db", "rental", "list"}, time.Now())

	var objs []*Rental
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *RentalDB) Add(ctx context.Context, model *Rental) error {
	defer goa.MeasureSince([]string{"goa", "db", "rental", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Rental", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *RentalDB) Update(ctx context.Context, model *Rental) error {
	defer goa.MeasureSince([]string{"goa", "db", "rental", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Rental", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *RentalDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "rental", "delete"}, time.Now())

	var obj Rental

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Rental", "error", err.Error())
		return err
	}

	return nil
}
