package cqs_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lucianogarciaz/kit/cqs"
	"github.com/lucianogarciaz/kit/vo"
)

func TestEventHydrate(t *testing.T) {
	require := require.New(t)

	var e cqs.BasicEvent

	id := vo.NewID()
	aggRootID := vo.NewID()
	eventName := cqs.EventName("event_name")
	dt := vo.NewDateTime(time.Now())
	version := cqs.EventVersion(1)

	e.Hydrate(id, eventName, dt, aggRootID, version)

	require.Equal(id, e.EventID())
	require.Equal(aggRootID, e.EventAggregateRootID())
	require.Equal(eventName, e.EventName())
	require.Equal(dt, e.EventAt())
	require.Equal(version, e.EventVersion())
}
