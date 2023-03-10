package vo_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lucianogarciaz/kit/vo"
)

func TestDateTimeNow(t *testing.T) {
	dt := vo.DateTimeNow()
	require.Equal(t, vo.NewDateTime(time.Time(dt)), dt)
}

func TestDateTimeValue(t *testing.T) {
	require := require.New(t)

	t.Run(`Given a DateTime,
	when value is called,
	then it returns the value formated`, func(t *testing.T) {
		ts := time.Date(1993, 2, 17, 12, 45, 12, 0, time.UTC)

		dt := vo.NewDateTime(ts)

		v, err := dt.Value()
		require.NoError(err)
		require.Equal("1993-02-17T12:45:12Z", v)
	})
}

func TestDateTimeScan(t *testing.T) {
	require := require.New(t)

	t.Run(`Given a unexpected type,
	when scan is called,
	then it returns error`, func(t *testing.T) {
		var dt vo.DateTime
		require.Error(dt.Scan(1234))
	})

	t.Run(`Given a string without format date,
	when scan is called,
	then it returns error`, func(t *testing.T) {
		var dt vo.DateTime
		require.Error(dt.Scan("123"))
	})

	t.Run(`Given a date string,
	when scan is called,
	then it does not return error and fills DateTime`, func(t *testing.T) {
		ts := time.Date(1993, 2, 17, 12, 45, 12, 0, time.UTC)

		var dt vo.DateTime
		require.NoError(dt.Scan(ts.Format(time.RFC3339Nano)))
		require.Equal(ts, time.Time(dt))
	})

	t.Run(`Given a date time.Time type,
	when scan is called,
	then it does not return error and fills DateTime`, func(t *testing.T) {
		ts := time.Date(1993, 2, 17, 12, 45, 12, 0, time.UTC)

		var dt vo.DateTime
		require.NoError(dt.Scan(ts))
		require.Equal(ts, time.Time(dt))
	})
}

func TestDateTimeJSONMarshaling(t *testing.T) {
	require := require.New(t)

	var (
		dt  = vo.NewDateTime(time.Now())
		b   []byte
		err error
	)

	t.Run(`Given a date
	when it's marshaled,
	then it returns a byte's slice and no error`, func(t *testing.T) {
		b, err = json.Marshal(&dt)
		require.NoError(err)
		require.NotEmpty(b)
	})

	t.Run(`Given a marshaled date,
	when it's unmarshaled,
	then the it returns before marshaled datetime`, func(t *testing.T) {
		var unmarshaled vo.DateTime
		require.NoError(json.Unmarshal(b, &unmarshaled))
		require.Equal(time.Time(dt).Unix(), time.Time(unmarshaled).Unix())
	})

	t.Run(`Given a wrong marshaled date,
	when it's unmarshaled,
	then the it returns an error`, func(t *testing.T) {
		var unmarshaled vo.DateTime
		require.Error(json.Unmarshal(nil, &unmarshaled))
	})
}

func TestDateTimeEqual(t *testing.T) {
	require := require.New(t)
	now := time.Now()

	t.Run(`Given a DateTime value,
	when Equal is called with another DateTime with same value,
	then it returns true`, func(t *testing.T) {
		require.True(vo.NewDateTime(now).Equal(vo.NewDateTime(now)))
	})

	t.Run(`Given a DateTime value,
	when Equal is called with another DateTime with different value,
	then it returns false`, func(t *testing.T) {
		require.False(vo.NewDateTime(now).Equal(vo.NewDateTime(now.Add(10 * time.Second))))
	})
}

func TestDateTimeFormat(t *testing.T) {
	require := require.New(t)
	expectedTimeStr := "2021-01-25T01:00:00Z"
	testTime, err := time.Parse(time.RFC3339, expectedTimeStr)
	require.NoError(err)

	t.Run(`Given a DateTime value,
	when String is called with a standard layout string,
	then it returns a string representation of that DateTime`, func(t *testing.T) {
		layout := "2006-01-02T01:00:00Z"
		require.Equal(vo.NewDateTime(testTime).String(layout), expectedTimeStr)
	})
}

func TestDateTimeIsZero(t *testing.T) {
	require := require.New(t)

	t.Run(`Given a zero-valued DateTime,
	when IsZero is called on it,
	then it returns true`, func(t *testing.T) {
		require.True(vo.DateTime{}.IsZero())
	})

	t.Run(`Given a non zero-valued DateTime,
	when IsZero is called on it,
	then it returns false`, func(t *testing.T) {
		require.False(vo.NewDateTime(time.Now()).IsZero())
	})
}

func TestDateTimeHelpers(t *testing.T) {
	require := require.New(t)

	t.Run(`Given a not empty DateTime,
	when TimeRangePtr is called,
	then it returns an pointer to the passed DateTime`, func(t *testing.T) {
		dt := vo.NewDateTime(time.Now())
		require.NotNil(vo.DateTimePtr(dt))
	})

	t.Run(`Given an empty DateTime,
	when DateTimePtr is called,
	then nil pointer is returned`, func(t *testing.T) {
		dt := vo.DateTime{}
		require.Nil(vo.DateTimePtr(dt))
	})

	t.Run(`Given a function that returns DateTime value,
	when nil is passed,
	then returns zero value`, func(t *testing.T) {
		require.Equal(vo.DateTime{}, vo.DateTimeVal(nil))
	})

	t.Run(`Given a function that returns DateTime value,
	when DateTime pointer is passed,
	then returns zero value`, func(t *testing.T) {
		dt := vo.NewDateTime(time.Now())
		require.Equal(dt, vo.DateTimeVal(&dt))
	})
}
