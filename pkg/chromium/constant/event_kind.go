package constant

type EventKind string

const (
	EventKindCreated   EventKind = "created"
	EventKindDestroyed EventKind = "destroyed"
	EventKindChanged   EventKind = "changed"
	EventKindAttached  EventKind = "attached"
	EventKindDetached  EventKind = "detached"
)
