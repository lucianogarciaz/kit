package obs

import "context"

var _ Metrics = NoopMetrics{}

// NoopMetrics is a metrics implementation that does nothing. It's for testing purposes.
type NoopMetrics struct{}

// Count does nothing and returns always nil.
func (NoopMetrics) Count(context.Context, string, float64, ...Tag) error {
	return nil
}

// Gauge does nothing and returns always nil.
func (NoopMetrics) Gauge(context.Context, string, float64, ...Tag) error {
	return nil
}

// Histogram does nothing and returns always nil.
func (NoopMetrics) Histogram(context.Context, string, float64, ...Tag) error {
	return nil
}
