package face

import (
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/types/run_key"
)

type Latest interface {
	Owner() string
	Latest() (map[run_key.Key]*run.Run, error)
}
