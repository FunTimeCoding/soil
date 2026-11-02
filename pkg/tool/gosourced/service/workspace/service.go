package workspace

import "github.com/funtimecoding/soil/pkg/source/index/cache"

type Service struct {
	workspaces *cache.Cache
	full       bool
}
