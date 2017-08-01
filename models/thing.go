package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/pop/nulls"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Thing struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Title     nulls.String `json:"title" db:"title"`
	Content   string       `json:"content" db:"content"`
}

// String is not required by pop and may be deleted
func (t Thing) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Things is not required by pop and may be deleted
type Things []Thing

// String is not required by pop and may be deleted
func (t Things) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (t *Thing) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Content, Name: "Content"},
	), nil
}

// ValidateSave gets run every time you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (t *Thing) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (t *Thing) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
