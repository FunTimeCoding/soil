package face

import "go/types"

type Set struct {
	byMethod map[string][]*types.Interface
}
