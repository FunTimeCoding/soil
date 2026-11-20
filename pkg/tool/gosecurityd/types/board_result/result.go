package board_result

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"

type Result struct {
	Total      int
	Supported  int
	Tracked    int
	Untracked  int
	Stale      int
	Upgradable int
	Rows       []*system_result.Summary
}
