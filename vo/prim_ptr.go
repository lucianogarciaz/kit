package vo

import (
	"reflect"
	"time"
)

// BoolPtr returns a pointer to the string value passed in.
func BoolPtr(b bool) *bool {
	return &b
}

// BoolValue returns the value or false if the pointer is nil.
func BoolValue(b *bool) (v bool) {
	if b == nil {
		return v
	}

	return *b
}

// BoolInterface returns a pointer of the interface value.
func BoolInterface(i interface{}) *bool {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		b, ok := i.(*bool)
		if !ok {
			return nil
		}

		return b
	}

	b, ok := i.(bool)
	if !ok {
		return nil
	}

	return &b
}

// StringPtr returns a pointer to the string value passed in.
func StringPtr(v string) *string {
	return &v
}

// StringValue returns the value or "" if the pointer is nil.
func StringValue(s *string) (v string) {
	if s == nil {
		return v
	}

	return *s
}

// StringInterface returns a pointer of the interface value.
func StringInterface(i interface{}) *string {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		v, ok := i.(*string)
		if !ok {
			return nil
		}

		return v
	}

	v, ok := i.(string)
	if !ok {
		return nil
	}

	return &v
}

// IntPtr returns a pointer to the int value passed in.
func IntPtr(v int) *int {
	return &v
}

// IntValue returns the value or 0 if the pointer is nil.
func IntValue(i *int) (v int) {
	if i == nil {
		return v
	}

	return *i
}

// IntInterface returns a pointer of the interface value.
func IntInterface(i interface{}) *int {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		v, ok := i.(*int)
		if !ok {
			return nil
		}

		return v
	}

	v, ok := i.(int)
	if !ok {
		return nil
	}

	return &v
}

// Int32Ptr returns a pointer to the int value passed in.
func Int32Ptr(v int32) *int32 {
	return &v
}

// Int32Value returns the value or 0 if the pointer is nil.
func Int32Value(i *int32) (v int32) {
	if i == nil {
		return v
	}

	return *i
}

// Int32Interface returns a pointer of the interface value.
func Int32Interface(i interface{}) *int32 {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		v, ok := i.(*int32)
		if !ok {
			return nil
		}

		return v
	}

	v, ok := i.(int32)
	if !ok {
		return nil
	}

	return &v
}

// Int64Ptr returns a pointer to the int value passed in.
func Int64Ptr(v int64) *int64 {
	return &v
}

// Int64Value returns the value or 0 if the pointer is nil.
func Int64Value(i *int64) (v int64) {
	if i == nil {
		return v
	}

	return *i
}

// Int64Interface returns a pointer of the interface value.
func Int64Interface(i interface{}) *int64 {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		v, ok := i.(*int64)
		if !ok {
			return nil
		}

		return v
	}

	v, ok := i.(int64)
	if !ok {
		return nil
	}

	return &v
}

// Float32Ptr returns a pointer to the float32 value passed in.
func Float32Ptr(v float32) *float32 {
	return &v
}

// Float32Value returns the value or 0 if the pointer is nil.
func Float32Value(f *float32) (v float32) {
	if f == nil {
		return v
	}

	return *f
}

// Float32Interface returns a pointer of the interface value.
func Float32Interface(i interface{}) *float32 {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		v, ok := i.(*float32)
		if !ok {
			return nil
		}

		return v
	}

	v, ok := i.(float32)
	if !ok {
		return nil
	}

	return &v
}

// Float64Ptr returns a pointer to the float64 value passed in.
func Float64Ptr(v float64) *float64 {
	return &v
}

// Float64Value returns the value or 0 if the pointer is nil.
func Float64Value(f *float64) (v float64) {
	if f == nil {
		return v
	}

	return *f
}

// Float64Interface returns a pointer of the interface value.
func Float64Interface(i interface{}) *float64 {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		v, ok := i.(*float64)
		if !ok {
			return nil
		}

		return v
	}

	v, ok := i.(float64)
	if !ok {
		return nil
	}

	return &v
}

// TimePtr returns a pointer to the TimePtr value passed in.
func TimePtr(time time.Time) *time.Time {
	return &time
}

// TimeValue returns the value or zero date if the pointer is nil.
func TimeValue(t *time.Time) (v time.Time) {
	if t == nil {
		return v
	}

	return *t
}
