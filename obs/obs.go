package obs

import (
	"context"
	"fmt"
	"time"

	"github.com/lucianogarciaz/kit/cqs"
)

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

func CommandHandlerObsMiddleware[C cqs.Command](obs Observer) cqs.CommandHandlerMiddleware[C] {
	return func(h cqs.CommandHandler[C]) cqs.CommandHandler[C] {
		return cqs.CommandHandlerFunc[C](func(ctx context.Context, cmd C) ([]cqs.Event, error) {
			defer func(begin time.Time) {
				elapsed := time.Since(begin)
				_ = obs.Log(LevelInfo, fmt.Sprintf("command: %s with latency: %f", cmd.CommandName(), elapsed.Seconds()))
			}(time.Now())

			events, err := h.Handle(ctx, cmd)
			if err != nil {
				_ = obs.Log(LevelError, fmt.Sprintf("command: %s with error: %s", cmd.CommandName(), err.Error()))
			}

			return events, err
		})
	}
}

// QueryHandlerObsMiddleware is a middleware function to make query handlers observables.
func QueryHandlerObsMiddleware[Q cqs.Query, R cqs.QueryResult](obs Observer) cqs.QueryHandlerMiddleware[Q, R] {
	return func(h cqs.QueryHandler[Q, R]) cqs.QueryHandler[Q, R] {
		return cqs.QueryHandlerFunc[Q, R](func(ctx context.Context, query Q) (R, error) {
			defer func(begin time.Time) {
				elapsed := time.Since(begin)
				_ = obs.Log(LevelInfo, fmt.Sprintf("query: %s, latency: %f",
					query.QueryName(),
					elapsed.Seconds()),
				)
			}(time.Now())
			result, err := h.Handle(ctx, query)
			if err != nil {
				_ = obs.Log(LevelError, fmt.Sprintf("query: %s, error: %s", query.QueryName(), err.Error()))
			}

			return result, err
		})
	}
}
