package unit

import "time"

func fired(c chan struct{}) bool {
	select {
	case <-c:
		return true
	case <-time.After(time.Second):
		return false
	}
}
