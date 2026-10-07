package bookmark

import "github.com/funtimecoding/soil/pkg/brave/bookmark/file"

type Root struct {
	Bar    *file.Node `json:"bookmark_bar"`
	Other  *file.Node `json:"other"`
	Synced *file.Node `json:"synced"`
}
