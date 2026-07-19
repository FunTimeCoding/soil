package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/jackc/pgconn"
	"testing"
)

func TestIsErrorCode(t *testing.T) {
	assert.True(t, relational.IsErrorCode(&pgconn.PgError{Code: "a"}, "a"))
	assert.False(t, relational.IsErrorCode(&pgconn.PgError{Code: "b"}, "a"))
}
