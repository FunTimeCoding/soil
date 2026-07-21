package unit_test

import (
	"charm.land/bubbles/v2/table"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/example_list"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/example_table"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/tick"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/toast"
	"github.com/funtimecoding/soil/pkg/bubbletea/table/example_country"
	"github.com/funtimecoding/soil/pkg/bubbletea/table/item"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestConstructors(t *testing.T) {
	assert.NotNil(t, example_list.New())
	assert.NotNil(t, example_table.New(&table.Model{}))
	assert.NotNil(t, fetch.Command())
	assert.NotNil(t, monitor.New(false))
	assert.NotNil(t, tick.Command())
	assert.NotNil(t, toast.New(0, upper.Alfa))
	assert.NotNil(t, example_country.New())
	assert.NotNil(t, item.New(true))
}
