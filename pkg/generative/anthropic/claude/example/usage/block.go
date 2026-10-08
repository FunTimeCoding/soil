package usage

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/example/common"
	"time"
)

type Block struct {
	start   time.Time
	end     time.Time
	entries []*common.Timestamped
}
