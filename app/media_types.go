// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "forent": Application Media Types
//
// Command:
// $ goagen
// --design=forent_api/design
// --out=$(GOPATH)/src/forent_api
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// article (default view)
//
// Identifier: application/vnd.article+json; view=default
type Article struct {
	// body 1
	Body1 string `form:"body1" json:"body1" xml:"body1"`
	// body 2
	Body2 string `form:"body2" json:"body2" xml:"body2"`
	// body 3
	Body3 string `form:"body3" json:"body3" xml:"body3"`
	// body 4
	Body4 string `form:"body4" json:"body4" xml:"body4"`
	// body 5
	Body5 string `form:"body5" json:"body5" xml:"body5"`
	// category ID
	CategoryID int `form:"category_id" json:"category_id" xml:"category_id"`
	// Unique article ID
	ID int `form:"id" json:"id" xml:"id"`
	// item image 1
	Image1 string `form:"image1" json:"image1" xml:"image1"`
	// item image 2
	Image2 string `form:"image2" json:"image2" xml:"image2"`
	// item image 3
	Image3 string `form:"image3" json:"image3" xml:"image3"`
	// item image 4
	Image4 string `form:"image4" json:"image4" xml:"image4"`
	// item image 5
	Image5 string `form:"image5" json:"image5" xml:"image5"`
	// introduction
	Introduction string `form:"introduction" json:"introduction" xml:"introduction"`
	// item description
	ItemDescription string `form:"item_description" json:"item_description" xml:"item_description"`
	// item ID
	ItemID int `form:"item_id" json:"item_id" xml:"item_id"`
	// title 1
	Title1 string `form:"title1" json:"title1" xml:"title1"`
	// title 2
	Title2 string `form:"title2" json:"title2" xml:"title2"`
	// title 3
	Title3 string `form:"title3" json:"title3" xml:"title3"`
	// title 4
	Title4 string `form:"title4" json:"title4" xml:"title4"`
	// title 5
	Title5 string `form:"title5" json:"title5" xml:"title5"`
	// user ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the Article media type instance.
func (mt *Article) Validate() (err error) {

	if mt.Title1 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title1"))
	}
	if mt.Title2 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title2"))
	}
	if mt.Title3 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title3"))
	}
	if mt.Title4 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title4"))
	}
	if mt.Title5 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title5"))
	}
	if mt.Body1 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body1"))
	}
	if mt.Body2 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body2"))
	}
	if mt.Body3 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body3"))
	}
	if mt.Body4 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body4"))
	}
	if mt.Body5 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body5"))
	}
	if mt.Introduction == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "introduction"))
	}
	if mt.ItemDescription == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "item_description"))
	}
	if mt.Image1 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image1"))
	}
	if mt.Image2 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image2"))
	}
	if mt.Image3 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image3"))
	}
	if mt.Image4 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image4"))
	}
	if mt.Image5 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image5"))
	}
	return
}

// ArticleCollection is the media type for an array of Article (default view)
//
// Identifier: application/vnd.article+json; type=collection; view=default
type ArticleCollection []*Article

// Validate validates the ArticleCollection media type instance.
func (mt ArticleCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Authentication (default view)
//
// Identifier: application/vnd.authentication+json; view=default
type Authentication struct {
	// address flag
	Email bool `form:"email" json:"email" xml:"email"`
	// Unique auth ID
	ID int `form:"id" json:"id" xml:"id"`
	// identification flag
	Identification bool `form:"identification" json:"identification" xml:"identification"`
	// phone flag
	Phone bool `form:"phone" json:"phone" xml:"phone"`
}

// Validate validates the Authentication media type instance.
func (mt *Authentication) Validate() (err error) {

	return
}

// Category (default view)
//
// Identifier: application/vnd.category+json; view=default
type Category struct {
	// Unique category ID
	ID int `form:"id" json:"id" xml:"id"`
	// middle category id
	MiddleCategoryID int `form:"middle_category_id" json:"middle_category_id" xml:"middle_category_id"`
	// category name
	Name string `form:"name" json:"name" xml:"name"`
	// category kana name
	NameBase string `form:"name_base" json:"name_base" xml:"name_base"`
	// category english name
	NameEn string `form:"name_en" json:"name_en" xml:"name_en"`
}

// Validate validates the Category media type instance.
func (mt *Category) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.NameBase == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name_base"))
	}
	if mt.NameEn == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name_en"))
	}
	return
}

// CategoryCollection is the media type for an array of Category (default view)
//
// Identifier: application/vnd.category+json; type=collection; view=default
type CategoryCollection []*Category

// Validate validates the CategoryCollection media type instance.
func (mt CategoryCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Comment (default view)
//
// Identifier: application/vnd.comment+json; view=default
type Comment struct {
	// Unique comment ID
	ID int `form:"id" json:"id" xml:"id"`
	// item id
	ItemID int `form:"item_id" json:"item_id" xml:"item_id"`
	// comment_id or null
	ReplyTo int `form:"reply_to" json:"reply_to" xml:"reply_to"`
	// comment text
	Text string `form:"text" json:"text" xml:"text"`
	// comment user id
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the Comment media type instance.
func (mt *Comment) Validate() (err error) {

	if mt.Text == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "text"))
	}

	return
}

// CommentCollection is the media type for an array of Comment (default view)
//
// Identifier: application/vnd.comment+json; type=collection; view=default
type CommentCollection []*Comment

// Validate validates the CommentCollection media type instance.
func (mt CommentCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Impression (default view)
//
// Identifier: application/vnd.impression+json; view=default
type Impression struct {
	// Unique impression ID
	ID int `form:"id" json:"id" xml:"id"`
	// reviewee user id
	RevieweeID int `form:"reviewee_id" json:"reviewee_id" xml:"reviewee_id"`
	// reviewer user id
	ReviewerID int `form:"reviewer_id" json:"reviewer_id" xml:"reviewer_id"`
	// impression score
	Score int `form:"score" json:"score" xml:"score"`
	// impression text
	Text string `form:"text" json:"text" xml:"text"`
}

// Validate validates the Impression media type instance.
func (mt *Impression) Validate() (err error) {

	if mt.Text == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "text"))
	}

	return
}

// A item of user (default view)
//
// Identifier: application/vnd.item+json; view=default
type Item struct {
	// Category ID
	CategoryID int `form:"category_id" json:"category_id" xml:"category_id"`
	// compensation of item
	Compensation int `form:"compensation" json:"compensation" xml:"compensation"`
	// description of item
	Description string `form:"description" json:"description" xml:"description"`
	// Unique item ID
	ID int `form:"id" json:"id" xml:"id"`
	// item image 1
	Image1 string `form:"image1" json:"image1" xml:"image1"`
	// item image 2
	Image2 string `form:"image2" json:"image2" xml:"image2"`
	// item image 3
	Image3 string `form:"image3" json:"image3" xml:"image3"`
	// item image 4
	Image4 string `form:"image4" json:"image4" xml:"image4"`
	// Name of item
	Name string `form:"name" json:"name" xml:"name"`
	// Place ID
	PlaceID int `form:"place_id" json:"place_id" xml:"place_id"`
	// price of item
	Price int `form:"price" json:"price" xml:"price"`
	// user ID
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the Item media type instance.
func (mt *Item) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}

	if mt.Image1 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image1"))
	}
	if mt.Image2 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image2"))
	}
	if mt.Image3 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image3"))
	}
	if mt.Image4 == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image4"))
	}
	return
}

// ItemCollection is the media type for an array of Item (default view)
//
// Identifier: application/vnd.item+json; type=collection; view=default
type ItemCollection []*Item

// Validate validates the ItemCollection media type instance.
func (mt ItemCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Large Category (default view)
//
// Identifier: application/vnd.largecategory+json; view=default
type Largecategory struct {
	// Unique large category ID
	ID int `form:"id" json:"id" xml:"id"`
	// category name
	Name string `form:"name" json:"name" xml:"name"`
	// category kana name
	NameBase string `form:"name_base" json:"name_base" xml:"name_base"`
	// category english name
	NameEn string `form:"name_en" json:"name_en" xml:"name_en"`
}

// Validate validates the Largecategory media type instance.
func (mt *Largecategory) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.NameBase == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name_base"))
	}
	if mt.NameEn == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name_en"))
	}
	return
}

// LargecategoryCollection is the media type for an array of Largecategory (default view)
//
// Identifier: application/vnd.largecategory+json; type=collection; view=default
type LargecategoryCollection []*Largecategory

// Validate validates the LargecategoryCollection media type instance.
func (mt LargecategoryCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Middle Category (default view)
//
// Identifier: application/vnd.middlecategory+json; view=default
type Middlecategory struct {
	// Unique middle category ID
	ID int `form:"id" json:"id" xml:"id"`
	// large category id
	LargeCategoryID int `form:"large_category_id" json:"large_category_id" xml:"large_category_id"`
	// category name
	Name string `form:"name" json:"name" xml:"name"`
	// category kana name
	NameBase string `form:"name_base" json:"name_base" xml:"name_base"`
	// category english name
	NameEn string `form:"name_en" json:"name_en" xml:"name_en"`
}

// Validate validates the Middlecategory media type instance.
func (mt *Middlecategory) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.NameBase == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name_base"))
	}
	if mt.NameEn == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name_en"))
	}
	return
}

// MiddlecategoryCollection is the media type for an array of Middlecategory (default view)
//
// Identifier: application/vnd.middlecategory+json; type=collection; view=default
type MiddlecategoryCollection []*Middlecategory

// Validate validates the MiddlecategoryCollection media type instance.
func (mt MiddlecategoryCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Offer (default view)
//
// Identifier: application/vnd.offer+json; view=default
type Offer struct {
	// offer accept
	Accepted bool `form:"accepted" json:"accepted" xml:"accepted"`
	// rental end at
	EndAt time.Time `form:"end_at" json:"end_at" xml:"end_at"`
	// Unique offer ID
	ID int `form:"id" json:"id" xml:"id"`
	// item id
	ItemID int `form:"item_id" json:"item_id" xml:"item_id"`
	// item id
	OwnerID int `form:"owner_id" json:"owner_id" xml:"owner_id"`
	// offer price
	Price int `form:"price" json:"price" xml:"price"`
	// rental start at
	StartAt time.Time `form:"start_at" json:"start_at" xml:"start_at"`
	// offer user id
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the Offer media type instance.
func (mt *Offer) Validate() (err error) {

	return
}

// Place (default view)
//
// Identifier: application/vnd.place+json; view=default
type Place struct {
	// google place id
	GooglePlaceID string `form:"google_place_id" json:"google_place_id" xml:"google_place_id"`
	// Unique place ID
	ID int `form:"id" json:"id" xml:"id"`
	// place latitude
	Latitude float64 `form:"latitude" json:"latitude" xml:"latitude"`
	// place longitude
	Longitude float64 `form:"longitude" json:"longitude" xml:"longitude"`
	// place name
	Name string `form:"name" json:"name" xml:"name"`
	// receive type
	Type int `form:"type" json:"type" xml:"type"`
}

// Validate validates the Place media type instance.
func (mt *Place) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if mt.GooglePlaceID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "google_place_id"))
	}

	return
}

// profile (default view)
//
// Identifier: application/vnd.profile+json; view=default
type Profile struct {
	// address
	Address string `form:"address" json:"address" xml:"address"`
	// avatar image url
	AvatarImage string `form:"avatar_image" json:"avatar_image" xml:"avatar_image"`
	// cover image url
	CoverImage string `form:"cover_image" json:"cover_image" xml:"cover_image"`
	// Unique profile ID
	ID int `form:"id" json:"id" xml:"id"`
	// user introduciton
	Introduction string `form:"introduction" json:"introduction" xml:"introduction"`
	// phone number
	Phone int `form:"phone" json:"phone" xml:"phone"`
}

// Validate validates the Profile media type instance.
func (mt *Profile) Validate() (err error) {

	if mt.Introduction == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "introduction"))
	}
	if mt.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}

	if mt.AvatarImage == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "avatar_image"))
	}
	if mt.CoverImage == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "cover_image"))
	}
	return
}

// Offer (default view)
//
// Identifier: application/vnd.rental+json; view=default
type Rental struct {
	// delivered at
	DeliveredAt time.Time `form:"delivered_at" json:"delivered_at" xml:"delivered_at"`
	// Unique rental ID
	ID int `form:"id" json:"id" xml:"id"`
	// offer id
	OfferID int `form:"offer_id" json:"offer_id" xml:"offer_id"`
	// returned at
	ReturnedAt time.Time `form:"returned_at" json:"returned_at" xml:"returned_at"`
}

// Validate validates the Rental media type instance.
func (mt *Rental) Validate() (err error) {

	return
}

// Review (default view)
//
// Identifier: application/vnd.review+json; view=default
type Review struct {
	// Unique review ID
	ID int `form:"id" json:"id" xml:"id"`
	// image
	Image string `form:"image" json:"image" xml:"image"`
	// item id
	ItemID int `form:"item_id" json:"item_id" xml:"item_id"`
	// score
	Score int `form:"score" json:"score" xml:"score"`
	// review text
	Text string `form:"text" json:"text" xml:"text"`
	// review user id
	UserID int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the Review media type instance.
func (mt *Review) Validate() (err error) {

	if mt.Text == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "text"))
	}

	if mt.Image == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image"))
	}
	return
}

// user (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// email
	Email string `form:"email" json:"email" xml:"email"`
	// first name
	FirstName string `form:"first_name" json:"first_name" xml:"first_name"`
	// Unique user ID
	ID int `form:"id" json:"id" xml:"id"`
	// last_name
	LastName string `form:"last_name" json:"last_name" xml:"last_name"`
	// password
	Password string `form:"password" json:"password" xml:"password"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if mt.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	return
}
