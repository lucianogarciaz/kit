package vo_test

import (
	"testing"

	"github.com/lucianogarciaz/kit/vo"
	"github.com/stretchr/testify/require"
)

func TestParseID(t *testing.T) {
	require := require.New(t)
	id := vo.NewID()
	parsedID, err := vo.ParseID(id.String())
	require.NoError(err)
	require.IsType(vo.ID{}, parsedID)
}

func TestUnmarshalBinary(t *testing.T) {
	require := require.New(t)
	id := vo.NewID()
	marshaledID, err := id.MarshalBinary()
	require.NoError(err)
	var unmarshaledID vo.ID
	err = unmarshaledID.UnmarshalBinary(marshaledID)
	require.NoError(err)
	require.Equal(id, unmarshaledID)
}

func TestMustParseID(t *testing.T) {
	require := require.New(t)
	t.Run(`given a string id that isn't an UUID,
		when calling MustParseID,
		then it should panic`, func(t *testing.T) {
		id := "should panic"
		require.Panics(func() { vo.MustParseID(id) })
	})
	t.Run(`given a string that is an UUID,
		when calling MustParseID,
		then it should not panic and an ID must be returned`, func(t *testing.T) {
		stringID := vo.NewID().String()
		require.NotPanics(func() {
			id := vo.MustParseID(stringID)
			require.IsType(vo.ID{}, id)
		})
	})
}

func TestScan(t *testing.T) {
	require := require.New(t)
	t.Run(`Given a string that is an UUID,
		when calling to Scan,
		then it should return an id`, func(t *testing.T) {
		stringID := vo.NewID().String()
		var id vo.ID
		err := id.Scan(stringID)
		require.NoError(err)
		require.Equal(stringID, id.String())
	})
	t.Run(`Given an string that isn't an UUID,
		when calling to Scan,
		then it should return an error`, func(t *testing.T) {
		var id vo.ID
		err := id.Scan("some incorrect uuid")
		require.Error(err)
	})
	t.Run(`Given an array of bytes,
		when calling to Scan,
		then it should scan it and return no error`, func(t *testing.T) {
		expectedID := vo.NewID()
		bytesID, err := expectedID.MarshalBinary()
		require.NoError(err)
		var id vo.ID
		err = id.Scan(bytesID)
		require.NoError(err)
		require.Equal(expectedID, id)
	})
}

func TestValue(t *testing.T) {
	require := require.New(t)

	id := vo.NewID()
	value, err := id.Value()

	require.NoError(err)
	require.Equal(id.String(), value)
}

func TestIsEmpty(t *testing.T) {
	require := require.New(t)
	t.Run(`Given a non empty id,
		when calling to IsEmpty,
		then it should return false`, func(t *testing.T) {
		id := vo.NewID()
		require.False(id.IsEmpty())
	})
	t.Run(`Given an empty id,
		when calling to IsEmpty,
		then it should return true`, func(t *testing.T) {
		id := vo.ID{}
		require.True(id.IsEmpty())
	})
}

func TestIDPtr(t *testing.T) {
	ID := vo.NewID()
	require.IsType(t, &ID, vo.IDPtr(ID))
	require.Equal(t, &ID, vo.IDPtr(ID))
}

func TestIDValue(t *testing.T) {
	t.Run(`Given a ID pointer,
	when IDValue is called,
	then it returns the value`, func(t *testing.T) {
		v := vo.NewID()
		require.Equal(t, v, vo.IDValue(vo.IDPtr(v)))
	})

	t.Run(`Given a nil,
	when IDValue is called,
	then it returns the zero value`, func(t *testing.T) {
		require.Equal(t, vo.ID{}, vo.IDValue(nil))
	})
}
