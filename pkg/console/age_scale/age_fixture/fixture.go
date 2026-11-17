package age_fixture

import (
	"github.com/funtimecoding/soil/pkg/face"
	"time"
)

type Fixture struct {
	value    time.Duration
	function face.SprintFunction
}
