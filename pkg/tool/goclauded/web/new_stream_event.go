package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	"time"
)

func newStreamEvent(v event.Event) *stream_event.Event {
	result := stream_event.New()
	result.Identifier = v.Identifier
	result.SessionIdentifier = v.SessionIdentifier
	result.Kind = v.Kind
	result.Actor = v.Actor
	result.Created = v.CreatedAt.Format(time.RFC3339)
	result.Metadata = v.Metadata

	return result
}
