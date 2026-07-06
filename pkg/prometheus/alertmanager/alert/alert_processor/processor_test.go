package alert_processor

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestProcessor(t *testing.T) {
	assert.NotNil(t, New(nil, nil, nil, nil, nil, nil))
}
