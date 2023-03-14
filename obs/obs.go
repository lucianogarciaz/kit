package obs

import "context"

// Observer provides Metrics and Logger features.
type Observer interface {
	Metrics
	Logger
}

// NewObserver returns a simple observer.
func NewObserver(m Metrics, l Logger) Observer {
	return struct {
		Metrics
		Logger
	}{m, l}
}

// Metrics defines the possible values to measure.
type Metrics interface {
	Count(ctx context.Context, name string, value float64, tags ...Tag) error
	Gauge(ctx context.Context, name string, value float64, tags ...Tag) error
	Histogram(ctx context.Context, name string, value float64, tags ...Tag) error
}

type Logger interface {
	Log(level LogLevel, message string, payload ...PayloadEntry) error
}

// Tag is a key-value attribute to give context to a metric value.
type Tag struct {
	Name  string
	Value string
}

// PayloadEntry is a log entry payload.
type PayloadEntry interface{}
