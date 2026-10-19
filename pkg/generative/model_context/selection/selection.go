package selection

import "sync"

type Selection struct {
	sessions sync.Map
}
