package blessed

import "example/fakegorm"

func New() string {
	return fakegorm.Open("allowed")
}
