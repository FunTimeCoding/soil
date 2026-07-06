package relational

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"gorm.io/gorm"
	"testing"
)

func TestNotFound(t *testing.T) {
	assert.False(t, NotFound(nil))
	assert.False(t, NotFound(errors.New("test")))
	assert.True(t, NotFound(gorm.ErrRecordNotFound))
}
