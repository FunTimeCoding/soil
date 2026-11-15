package clock_tester

import (
	"sync"
	"time"
)

type Clock struct {
	lock sync.Mutex
	now  time.Time
}
