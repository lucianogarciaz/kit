package vo

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
)

type ID uuid.UUID

// NewID returns an ID with a UUID v4 value.
func NewID() ID {
	return ID(uuid.New())
}

// ParseID is self-described.
func ParseID(s string) (ID, error) {
	uuid, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}

	return ID(uuid), nil
}

// MustParseID parses a string and if has any errors will panic.
func MustParseID(s string) ID {
	id, err := ParseID(s)
	if err != nil {
		panic(fmt.Sprintf("ID %s parse: %s", s, err.Error()))
	}

	return id
}

// String is self-described.
func (id ID) String() string {
	return uuid.UUID(id).String()
}

// IsEmpty is self-described.
func (id ID) IsEmpty() bool {
	return id == ID{}
}

// Value returns the ID as a string to be stored in the database as uuid.
func (id ID) Value() (driver.Value, error) {
	return uuid.UUID(id).Value()
}

// Scan binds the value from the database with the type ID.
func (id *ID) Scan(src interface{}) error {
	var uuid uuid.UUID

	err := (uuid).Scan(src)
	if err != nil {
		return err
	}

	*id = ID(uuid)

	return nil
}

// MarshalBinary is self-described.
func (id ID) MarshalBinary() ([]byte, error) {
	return uuid.UUID(id).MarshalBinary()
}

// UnmarshalBinary is self-described.
func (id *ID) UnmarshalBinary(data []byte) error {
	var uuid uuid.UUID

	err := uuid.UnmarshalBinary(data)
	if err == nil {
		*id = ID(uuid)
	}

	return err
}

// IDPtr returns a pointer to the ID value passed in.
func IDPtr(v ID) *ID {
	return &v
}

// IDValue returns the value or zero ID if the pointer is nil.
func IDValue(id *ID) (v ID) {
	if id == nil {
		return v
	}

	return *id
}
