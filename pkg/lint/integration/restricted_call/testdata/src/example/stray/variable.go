package stray

import "example/fakegorm"

func Variable() string {
	return fakegorm.Shared
}

func Field(h fakegorm.Holder) string {
	return h.Shared
}

func Method(h fakegorm.Holder) string {
	return h.Open("method")
}
