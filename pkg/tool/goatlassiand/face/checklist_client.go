package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

type ChecklistClient interface {
	GetChecklist(
		x context.Context,
		key string,
	) ([]*client.ChecklistItem, error)
	AddChecklistItem(
		x context.Context,
		key string,
		text string,
	) ([]*client.ChecklistItem, error)
	EditChecklistItem(
		x context.Context,
		key string,
		index int,
		text string,
	) ([]*client.ChecklistItem, error)
	DeleteChecklistItem(
		x context.Context,
		key string,
		index int,
	) ([]*client.ChecklistItem, error)
	ToggleChecklistItem(
		x context.Context,
		key string,
		index int,
	) ([]*client.ChecklistItem, error)
}
