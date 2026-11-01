package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"testing"
)

func TestCloseClosesTheQueryClient(t *testing.T) {
	o := service_tester.New(t)
	o.Service.Close()
	o.Service.Close()
	assert.Integer(t, 1, o.Indexer.Closed)
}
