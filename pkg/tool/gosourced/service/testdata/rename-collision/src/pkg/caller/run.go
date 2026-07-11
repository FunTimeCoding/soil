package caller

import "example/pkg/outer/store"

func depot() string {
	return "alfa"
}

func Run() string {
	v := store.Store{}
	_ = v

	return depot()
}
