package vo_test

import (
	"testing"
	"time"

	"github.com/lucianogarciaz/kit/vo"

	"github.com/stretchr/testify/require"
)

func TestBoolPtr(t *testing.T) {
	in := true
	require.IsType(t, &in, vo.BoolPtr(in))
	require.Equal(t, &in, vo.BoolPtr(in))
}

func TestBoolValue(t *testing.T) {
	t.Run(`Given a bool pointer,
	when BoolValue is called,
	then it returns the value`, func(t *testing.T) {
		require.True(t, vo.BoolValue(vo.BoolPtr(true)))
	})

	t.Run(`Given a nil,
	when BoolValue is called,
	then it returns the zero value`, func(t *testing.T) {
		require.False(t, vo.BoolValue(nil))
	})
}

func TestBoolInterface(t *testing.T) {
	type testCase struct {
		in       interface{}
		expected *bool
	}

	for _, tc := range []testCase{
		{in: nil, expected: nil},
		{in: "string", expected: nil},
		{in: true, expected: vo.BoolPtr(true)},
		{in: false, expected: vo.BoolPtr(false)},
		{in: vo.StringPtr("string"), expected: nil},
		{in: vo.BoolPtr(true), expected: vo.BoolPtr(true)},
		{in: vo.BoolPtr(false), expected: vo.BoolPtr(false)},
	} {
		require.Exactly(t, tc.expected, vo.BoolInterface(tc.in))
	}
}

func TestStringValue(t *testing.T) {
	t.Run(`Given a string pointer,
	when StringValue is called,
	then it returns the value`, func(t *testing.T) {
		v := "hello"
		require.Equal(t, v, vo.StringValue(vo.StringPtr(v)))
	})

	t.Run(`Given a nil,
	when StringValue is called,
	then it returns the zero value`, func(t *testing.T) {
		require.Empty(t, vo.StringValue(nil))
	})
}

func TestStringInterface(t *testing.T) {
	type testCase struct {
		in       interface{}
		expected *string
	}

	for _, tc := range []testCase{
		{in: nil, expected: nil},
		{in: 55, expected: nil},
		{in: "string", expected: vo.StringPtr("string")},
		{in: "", expected: vo.StringPtr("")},
		{in: vo.BoolPtr(false), expected: nil},
		{in: vo.StringPtr(""), expected: vo.StringPtr("")},
		{in: vo.StringPtr("string"), expected: vo.StringPtr("string")},
	} {
		require.Exactly(t, tc.expected, vo.StringInterface(tc.in))
	}
}

func TestIntValue(t *testing.T) {
	t.Run(`Given an int pointer,
	when IntValue is called,
	then it returns the value`, func(t *testing.T) {
		v := 123
		require.Equal(t, v, vo.IntValue(vo.IntPtr(v)))
	})

	t.Run(`Given a nil,
	when IntValue is called,
	then it returns the zero value`, func(t *testing.T) {
		require.Equal(t, 0, vo.IntValue(nil))
	})
}

func TestIntInterface(t *testing.T) {
	type testCase struct {
		in       interface{}
		expected *int
	}

	for _, tc := range []testCase{
		{in: nil, expected: nil},
		{in: "string", expected: nil},
		{in: vo.StringPtr("string"), expected: nil},
		{in: 55, expected: vo.IntPtr(55)},
		{in: vo.IntPtr(55), expected: vo.IntPtr(55)},
	} {
		require.Exactly(t, tc.expected, vo.IntInterface(tc.in))
	}
}

func TestInt32Value(t *testing.T) {
	t.Run(`Given an int pointer,
	when Int32Value is called,
	then it returns the value`, func(t *testing.T) {
		var v int32 = 123
		require.Equal(t, v, vo.Int32Value(vo.Int32Ptr(v)))
	})

	t.Run(`Given a nil,
	when Int32Value is called,
	then it returns the zero value`, func(t *testing.T) {
		require.Equal(t, int32(0), vo.Int32Value(nil))
	})
}

func TestInt32Interface(t *testing.T) {
	type testCase struct {
		in       interface{}
		expected *int32
	}

	for _, tc := range []testCase{
		{in: nil, expected: nil},
		{in: "string", expected: nil},
		{in: vo.StringPtr("string"), expected: nil},
		{in: int32(55), expected: vo.Int32Ptr(55)},
		{in: vo.Int32Ptr(55), expected: vo.Int32Ptr(55)},
	} {
		require.Exactly(t, tc.expected, vo.Int32Interface(tc.in))
	}
}

func TestInt64Ptr(t *testing.T) {
	var in int64 = 123
	require.IsType(t, &in, vo.Int64Ptr(in))
	require.Equal(t, &in, vo.Int64Ptr(in))
}

func TestInt64Value(t *testing.T) {
	t.Run(`Given an int pointer,
	when Int64Value is called,
	then it returns the value`, func(t *testing.T) {
		var v int64 = 123
		require.Equal(t, v, vo.Int64Value(vo.Int64Ptr(v)))
	})

	t.Run(`Given a nil,
	when Int64Value is called,
	then it returns the zero value`, func(t *testing.T) {
		require.Equal(t, int64(0), vo.Int64Value(nil))
	})
}

func TestInt64Interface(t *testing.T) {
	type testCase struct {
		in       interface{}
		expected *int64
	}

	for _, tc := range []testCase{
		{in: nil, expected: nil},
		{in: "string", expected: nil},
		{in: vo.StringPtr("string"), expected: nil},
		{in: int64(55), expected: vo.Int64Ptr(55)},
		{in: vo.Int64Ptr(55), expected: vo.Int64Ptr(55)},
	} {
		require.Exactly(t, tc.expected, vo.Int64Interface(tc.in))
	}
}

func TestFloat32Value(t *testing.T) {
	t.Run(`Given a float pointer,
	when Float32Value is called,
	then it returns the value`, func(t *testing.T) {
		var v float32 = 123
		require.Equal(t, v, vo.Float32Value(vo.Float32Ptr(v)))
	})

	t.Run(`Given a nil,
	when Float32Value is called,
	then it returns the zero value`, func(t *testing.T) {
		require.Equal(t, float32(0), vo.Float32Value(nil))
	})
}

func TestFloat32Interface(t *testing.T) {
	type testCase struct {
		in       interface{}
		expected *float32
	}

	for _, tc := range []testCase{
		{in: nil, expected: nil},
		{in: "string", expected: nil},
		{in: vo.StringPtr("string"), expected: nil},
		{in: float32(55), expected: vo.Float32Ptr(55)},
		{in: vo.Float32Ptr(55), expected: vo.Float32Ptr(55)},
	} {
		require.Exactly(t, tc.expected, vo.Float32Interface(tc.in))
	}
}

func TestFloat64Value(t *testing.T) {
	t.Run(`Given a float pointer,
	when Float64Value is called,
	then it returns the value`, func(t *testing.T) {
		var v float64 = 123
		require.Equal(t, v, vo.Float64Value(vo.Float64Ptr(v)))
	})

	t.Run(`Given a nil,
	when Float64Value is called,
	then it returns the zero value`, func(t *testing.T) {
		require.Equal(t, float64(0), vo.Float64Value(nil))
	})
}

func TestFloat64Interface(t *testing.T) {
	type testCase struct {
		in       interface{}
		expected *float64
	}

	for _, tc := range []testCase{
		{in: nil, expected: nil},
		{in: "string", expected: nil},
		{in: vo.StringPtr("string"), expected: nil},
		{in: float64(55), expected: vo.Float64Ptr(55)},
		{in: vo.Float64Ptr(55), expected: vo.Float64Ptr(55)},
	} {
		require.Exactly(t, tc.expected, vo.Float64Interface(tc.in))
	}
}

func TestTimeValue(t *testing.T) {
	t.Run(`Given a time pointer,
	when Float64Value is called,
	then it returns the value`, func(t *testing.T) {
		v := time.Now()
		require.Equal(t, v, vo.TimeValue(vo.TimePtr(v)))
	})

	t.Run(`Given a nil,
	when Float64Value is called,
	then it returns the zero value`, func(t *testing.T) {
		require.Equal(t, time.Time{}, vo.TimeValue(nil))
	})
}
