package entry

import (
	"github.com/funtimecoding/soil/pkg/crap/function"
	"github.com/funtimecoding/soil/pkg/crap/mutation"
)

type Entry struct {
	Function  *function.Function `json:"function"`
	Coverage  float64            `json:"coverage"`
	Score     float64            `json:"score"`
	Missing   bool               `json:"missing"`
	Untrusted bool               `json:"untrusted"`
	Killed    int                `json:"killed"`
	Lived     []*mutation.Mutant `json:"lived"`
	Previous  *float64           `json:"previous"`
}
