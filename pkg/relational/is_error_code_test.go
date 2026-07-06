package relational

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/jackc/pgconn"
	"testing"
)

func TestIsErrorCode(t *testing.T) {
	assert.True(t, IsErrorCode(&pgconn.PgError{Code: "a"}, "a"))
	assert.False(t, IsErrorCode(&pgconn.PgError{Code: "b"}, "a"))
}
