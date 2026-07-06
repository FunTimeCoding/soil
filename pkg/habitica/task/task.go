package task

import "github.com/funtimecoding/soil/pkg/habitica/checklist_item"

type Task struct {
	ID        string                 `json:"id"`
	Text      string                 `json:"text"`
	Type      string                 `json:"type"`
	Notes     string                 `json:"notes"`
	Tags      []string               `json:"tags"`
	Completed bool                   `json:"completed"`
	Checklist []*checklist_item.Item `json:"checklist"`
	Streak    int                    `json:"streak"`
	Value     float64                `json:"value"`
}
