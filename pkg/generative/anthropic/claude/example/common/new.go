package common

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/usage_entry"
	"time"
)

func New(
	t time.Time,
	entry *usage_entry.Entry,
) *Timestamped {
	return &Timestamped{
		Time:  t,
		Entry: entry,
	}
}
