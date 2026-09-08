package constant

type EventKind int

const (
	MessageEvent EventKind = iota
	ReactionAddedEvent
	ReactionRemovedEvent
)
