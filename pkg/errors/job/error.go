package job

import "fmt"

func (e *JobError) Error() string {
	return fmt.Sprintf("job %v (%s): %v", e.Identifier, e.Kind, e.Wrapped)
}
