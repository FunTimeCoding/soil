package stray

import "example/fakegorm"

func Direct() string {
	return fakegorm.Open("stray")
}
