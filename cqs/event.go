package cqs

import "github.com/lucianogarciaz/kit/vo"

// An Event describes a change that has happened in the system.
type Event interface {
	EventID() vo.ID
	EventName() EventName
	EventAt() vo.DateTime
	EventAggregateRootID() vo.ID
	EventVersion() EventVersion
}

// EventName is self-described.
type EventName string

// EventVersion is self-described.
type EventVersion int

// Payload is self-described.
type Payload interface{}

var _ Event = &BasicEvent{}

// BasicEvent is the minimal domain event struct.
type BasicEvent struct {
	ID              vo.ID        `json:"id,omitempty"`
	Name            EventName    `json:"name,omitempty"`
	At              vo.DateTime  `json:"at,omitempty"`
	Version         EventVersion `json:"version"`
	AggregateRootID vo.ID        `json:"aggregate_root_id,omitempty"`
}

func (b BasicEvent) EventID() vo.ID {
	return b.ID
}

func (b BasicEvent) EventName() EventName {
	return b.Name
}

func (b BasicEvent) EventAt() vo.DateTime {
	return b.At
}

func (b BasicEvent) EventVersion() EventVersion {
	return b.Version
}

func (b BasicEvent) EventAggregateRootID() vo.ID {
	return b.AggregateRootID
}

// Hydrate hydrates the instance.
func (b *BasicEvent) Hydrate(id vo.ID, name EventName, at vo.DateTime, aggRootID vo.ID, version EventVersion) {
	b.ID = id
	b.Name = name
	b.At = at
	b.AggregateRootID = aggRootID
	b.Version = version
}
