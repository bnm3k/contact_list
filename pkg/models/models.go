package models

import "errors"

import "time"

//ErrNoRecord indicates that for a given db query, the
//respective record not found
var ErrNoRecord = errors.New("models:no matching record found")

//Contact holds fields from a contact Row
type Contact struct {
	ID         int
	Name       string
	Phone      string
	Address    string
	Favourites string
}

//FullContact extends Contact struct to include
//fields for created_at and updated_at
type FullContact struct {
	Contact
	Created time.Time
	Updated time.Time
}
