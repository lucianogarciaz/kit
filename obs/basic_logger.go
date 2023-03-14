package obs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lucianogarciaz/kit/vo"
)

type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warning"
	LevelError LogLevel = "error"
)

// Marshaler is a log payload marshaler.
type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
}

// LoggerOpt is the common type of functions that set options on Logger construction.
type LoggerOpt func(l *BasicLogger)

// TimeFunc is the function to set logs time.
type TimeFunc func() vo.DateTime

var _ Logger = BasicLogger{}

// BasicLogger is a Logger.
type BasicLogger struct {
	marshaler Marshaler
	writer    io.Writer
	timeFunc  TimeFunc
}

// Entry is a log entry.
type Entry struct {
	Message string         `json:"message,omitempty"`
	Payload []PayloadEntry `json:"payload,omitempty"`
	Time    string         `json:"time,omitempty"`
	Level   LogLevel       `json:"log_level,omitempty"`
}

// Log is self-described.
func (bl BasicLogger) Log(level LogLevel, message string, payload ...PayloadEntry) error {
	b, err := bl.marshaler.Marshal(Entry{
		Level:   level,
		Message: message,
		Payload: payload,
		Time:    bl.timeFunc().Format(time.RFC3339),
	})
	if err != nil {
		return fmt.Errorf("info: %w", err)
	}

	_, err = fmt.Fprintf(bl.writer, "%s\n", string(b))

	return err
}

// NewBasicLogger is a constructor.
func NewBasicLogger(opts ...LoggerOpt) BasicLogger {
	l := BasicLogger{
		writer:    os.Stdout,
		timeFunc:  vo.DateTimeNow,
		marshaler: jsonMarshaler{},
	}
	for _, opt := range opts {
		opt(&l)
	}

	return l
}

// MarshalerOpt is an option that sets the Marshaler.
func MarshalerOpt(m Marshaler) LoggerOpt {
	return func(l *BasicLogger) {
		l.marshaler = m
	}
}

// WriterOpt is an option that sets the writer.
func WriterOpt(w io.Writer) LoggerOpt {
	return func(l *BasicLogger) {
		l.writer = w
	}
}

var _ Marshaler = &jsonMarshaler{}

type jsonMarshaler struct{}

func (j jsonMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
