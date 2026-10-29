package failure

import "sync"

type Failure struct {
	owner  string
	errors map[string]error
	mutex  sync.Mutex
}
