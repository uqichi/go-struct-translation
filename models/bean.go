package models

import (
	"encoding/json"

	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Bean struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	ProducingAreaID uuid.UUID `json:"producing_area_id" db:"producing_area_id"`
	Producer        string    `json:"producer" db:"producer"`
	Area            *Area     `json:"area" belongs_to:"area"`
}

// String is not required by pop and may be deleted
func (b Bean) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Beans is not required by pop and may be deleted
type Beans []Bean

// String is not required by pop and may be deleted
func (b Beans) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *Bean) Validate(tx *pop.Connection) (*validate.Errors, error) {
	verrs := validate.Validate(
		&validators.StringIsPresent{Field: b.Name, Name: "Name"},
		&validators.UUIDIsPresent{Field: b.ProducingAreaID, Name: "ProducingAreaID"},
	)
	return verrs, nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *Bean) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *Bean) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	verrs := validate.Validate(
		&validators.UUIDIsPresent{Field: b.ID, Name: "ID"},
	)
	return verrs, nil
}

// BeforeDestroy implements pop.BeforeDestroyable
func (b *Bean) BeforeDestroy(tx *pop.Connection) error {
	return errors.New("bean should not be deleted")
}
