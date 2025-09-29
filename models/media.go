package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Media is used by pop to map your media database table to your go code.
type Media struct {
	ID         uuid.UUID `json:"id" db:"id"`
	AdminID    uuid.UUID `json:"admin_id" db:"admin_id"`
	CategoryID int       `json:"category_id" db:"category_id"`
	Title      string    `json:"title" db:"title"`
	Img_Url    string    `json:"img_url" db:"img_url"`
	
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (m Media) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Media is not required by pop and may be deleted
type Medias []Media

// String is not required by pop and may be deleted
func (m Medias) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (m *Media) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (m *Media) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (m *Media) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
