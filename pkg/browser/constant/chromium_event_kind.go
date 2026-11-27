package constant

type ChromiumEventKind string

const (
	ChromiumEventKindCreated   ChromiumEventKind = "created"
	ChromiumEventKindDestroyed ChromiumEventKind = "destroyed"
	ChromiumEventKindChanged   ChromiumEventKind = "changed"
	ChromiumEventKindAttached  ChromiumEventKind = "attached"
	ChromiumEventKindDetached  ChromiumEventKind = "detached"
)
