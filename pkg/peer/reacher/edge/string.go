package edge

import "fmt"

func (e *Edge) String() string {
	if e.Down {
		return fmt.Sprintf("%s down: %s", e.Host, e.Reason)
	}

	return fmt.Sprintf("%s up after %s", e.Host, e.Duration)
}
