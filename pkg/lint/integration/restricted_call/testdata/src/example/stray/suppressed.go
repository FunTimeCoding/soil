package stray

import "example/fakegorm"

func Suppressed() string {
	return fakegorm.Open("tolerated") // goanalyze:ignore
}
