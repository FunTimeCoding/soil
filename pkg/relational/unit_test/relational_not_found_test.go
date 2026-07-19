package unit_test

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational"
	"gorm.io/gorm"
	"testing"
)

func TestNotFound(t *testing.T) {
	assert.False(t, relational.NotFound(nil))
	assert.False(t, relational.NotFound(errors.New("test")))
	assert.True(t, relational.NotFound(gorm.ErrRecordNotFound))
}
