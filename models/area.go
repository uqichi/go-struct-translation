package models

import (
	"encoding/json"

	"github.com/gobuffalo/validate/validators"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/pkg/errors"
)

type Area struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

// String is not required by pop and may be deleted
func (a Area) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Areas is not required by pop and may be deleted
type Areas []Area

// String is not required by pop and may be deleted
func (a Areas) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Area) Validate(tx *pop.Connection) (*validate.Errors, error) {
	verrs := validate.Validate(
		&validators.StringIsPresent{Field: a.Name, Name: "Name"},
	)
	return verrs, nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Area) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Area) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	verrs := validate.Validate(
		&validators.UUIDIsPresent{Field: a.ID, Name: "ID"},
	)
	return verrs, nil
}

// BeforeDestroy implements pop.BeforeDestroyable
func (a *Area) BeforeDestroy(tx *pop.Connection) error {
	return errors.New("area should not be deleted")
}
