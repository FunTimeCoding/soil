package caller

import (
	"example/pkg/outer/store"
	"example/pkg/outer/store/sub"
)

func Run() string {
	v := store.Store{}
	_ = v

	return sub.Name()
}
