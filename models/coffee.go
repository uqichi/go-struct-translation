package models

import (
	"encoding/json"

	"github.com/gobuffalo/validate/validators"

	"github.com/pkg/errors"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Coffee struct {
	ID     uuid.UUID `json:"id" db:"id"`
	Name   string    `json:"name" db:"name"`
	Price  int       `json:"price" db:"price"`
	Cost   int       `json:"cost" db:"cost"`
	Style  Style     `json:"style" db:"style"`
	BeanID uuid.UUID `json:"bean_id" db:"bean_id"`
	Bean   *Bean     `json:"bean" belongs_to:"bean"`
}

// String is not required by pop and may be deleted
func (c Coffee) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Coffees is not required by pop and may be deleted
type Coffees []Coffee

// String is not required by pop and may be deleted
func (c Coffees) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Coffee) Validate(tx *pop.Connection) (*validate.Errors, error) {
	verrs := validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
		&validators.IntIsPresent{Field: c.Price, Name: "Price"},
		&validators.IntIsPresent{Field: int(c.Style), Name: "Style"},
		&validators.UUIDIsPresent{Field: c.BeanID, Name: "BeanID"},
	)
	return verrs, nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Coffee) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Coffee) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	verrs := validate.Validate(
		&validators.UUIDIsPresent{Field: c.ID, Name: "ID"},
	)
	return verrs, nil
}

// BeforeDestroy implements pop.BeforeDestroyable
func (c *Coffee) BeforeDestroy(tx *pop.Connection) error {
	return errors.New("coffee should not be deleted")
}
