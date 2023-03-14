package obs_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lucianogarciaz/kit/obs"
)

func TestBasicLoggerInfo(t *testing.T) {
	require := require.New(t)

	t.Run(`Given a logger with a marshaler that returns an error,
	when Log is called,
	then it returns the error`, func(t *testing.T) {
		marshalErr := errors.New("marshaler error")
		m := &MarshalerMock{
			MarshalFunc: func(v interface{}) ([]byte, error) {
				return nil, marshalErr
			},
		}

		l := obs.NewBasicLogger(obs.MarshalerOpt(m))
		const msg = "a message"

		err := l.Log(obs.LevelInfo, msg)

		require.True(errors.Is(err, marshalErr))
		require.Len(m.MarshalCalls(), 1)
	})

	t.Run(`Given a logger with a JSON marshaler and a writer,
	when Log is called with a message,
	then the logger entry is written to the writer`, func(t *testing.T) {
		var w bytes.Buffer
		l := obs.NewBasicLogger(obs.WriterOpt(&w))
		expectedLogLevel := obs.LevelInfo
		const msg = "a message"
		require.NoError(l.Log(expectedLogLevel, msg))

		var entry obs.Entry
		require.NoError(json.Unmarshal(w.Bytes(), &entry))

		require.Equal(msg, entry.Message)
		require.Equal(expectedLogLevel, entry.Level)
		require.Empty(entry.Payload)
		require.NotEmpty(entry.Time)
	})

	t.Run(`Given a logger with a custom marshaler and a writer,
	when Log is called with a message,
	then the logger entry is written to the writer`, func(t *testing.T) {
		m := &MarshalerMock{
			MarshalFunc: func(v interface{}) ([]byte, error) {
				entry, _ := v.(obs.Entry)
				return []byte(entry.Message), nil
			},
		}
		var w bytes.Buffer
		l := obs.NewBasicLogger(obs.MarshalerOpt(m), obs.WriterOpt(&w))

		const msg = "a message"
		require.NoError(l.Log(obs.LevelInfo, msg))
		require.Equal(msg+"\n", w.String())
		require.Len(m.MarshalCalls(), 1)
	})
}
