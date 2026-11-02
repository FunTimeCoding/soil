package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set_record"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_record"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_template"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field_record"
	"github.com/funtimecoding/soil/pkg/netbox/custom_link"
	"github.com/funtimecoding/soil/pkg/netbox/export_template"
	"github.com/funtimecoding/soil/pkg/netbox/notification_group"
	"github.com/funtimecoding/soil/pkg/netbox/source"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestCustomizationEntities(t *testing.T) {
	assert.NotNil(t, choice_set.New(choice_set_record.New()))
	assert.NotNil(t, configuration_context.New(configuration_record.New()))
	assert.NotNil(t, configuration_template.New(&netbox.ConfigTemplate{}))
	assert.NotNil(t, custom_field.New(custom_field_record.New()))
	assert.NotNil(t, custom_link.New(&netbox.CustomLink{}))
	assert.NotNil(t, export_template.New(&netbox.ExportTemplate{}))
	assert.NotNil(t, notification_group.New(&netbox.NotificationGroup{}))
	assert.NotNil(t, source.New(&netbox.DataSource{}))
	assert.NotNil(t, tag.New(&netbox.Tag{}))
}
