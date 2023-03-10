package vo

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

var (
	_ sql.Scanner      = &DateTime{}
	_ driver.Valuer    = DateTime{}
	_ json.Unmarshaler = &DateTime{}
	_ json.Marshaler   = &DateTime{}
)

// DateTime is a time.Time type.
type DateTime time.Time

// NewDateTime is a constructor.
func NewDateTime(t time.Time) DateTime {
	return DateTime(t.UTC().Truncate(time.Microsecond))
}

// DateTimeNow generates a datetime with the current date.
func DateTimeNow() DateTime {
	return NewDateTime(time.Now())
}

// Value returns the DateTime value to be stored in database.
func (dt DateTime) Value() (driver.Value, error) {
	t := time.Time(dt).
		UTC().
		Truncate(time.Microsecond).
		Format(time.RFC3339Nano)

	return t, nil
}

// Scan binds the value from the database with the type Time.
func (dt *DateTime) Scan(src interface{}) error {
	var (
		t   time.Time
		err error
	)

	switch v := src.(type) {
	case string:
		t, err = time.Parse(time.RFC3339Nano, v)
		if err != nil {
			return err
		}

	case time.Time:
		t = v
	default:
		return errors.New("invalid source type")
	}

	*dt = NewDateTime(t)

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (dt DateTime) MarshalJSON() ([]byte, error) {
	return time.Time(dt).MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (dt *DateTime) UnmarshalJSON(bytes []byte) error {
	var t time.Time

	err := t.UnmarshalJSON(bytes)
	if err != nil {
		return err
	}

	*dt = NewDateTime(t)

	return nil
}

// Equal checks if the calling DateTime has the same value as input's.
func (dt DateTime) Equal(input DateTime) bool {
	return time.Time(dt).Equal(time.Time(input))
}

// Format returns a string representation of the calling DateTime.
func (dt DateTime) Format(layout string) string {
	return time.Time(dt).Format(layout)
}

// IsZero returns true if the DateTime on which it was called is a zero-valued DateTime.
func (dt DateTime) IsZero() bool {
	return time.Time(dt).IsZero()
}

// DateTimePtr returns the pointer of a date-time given.
func DateTimePtr(dt DateTime) *DateTime {
	if dt.IsZero() {
		return nil
	}

	return &dt
}

// DateTimeVal return the value of a date-time pointer.
func DateTimeVal(dt *DateTime) DateTime {
	if dt != nil {
		return *dt
	}

	return DateTime{}
}
