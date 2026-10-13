package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/inventory/face"
	"github.com/jackc/pgx/v5/pgxpool"
	"testing"
)

func TestConnectionPoolsAreClosedByTheSharedPool(t *testing.T) {
	var connection any = (*pgxpool.Pool)(nil)
	_, okay := connection.(face.Closer)
	assert.True(t, okay)
}
