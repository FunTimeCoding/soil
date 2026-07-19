package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/export_template"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestExportTemplate(t *testing.T) {
	assert.NotNil(t, export_template.New(&netbox.ExportTemplate{}))
}
