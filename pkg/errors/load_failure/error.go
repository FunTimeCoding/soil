package load_failure

import "fmt"

func (e *LoadFailureError) Error() string {
	return fmt.Sprintf("%s: %v", e.Package, e.Cause)
}
