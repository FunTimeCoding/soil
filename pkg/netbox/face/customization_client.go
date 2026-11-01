package face

import (
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_assignment"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_context"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_template"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/netbox/custom_link"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/export_template"
	"github.com/funtimecoding/soil/pkg/netbox/notification_group"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
)

type CustomizationClient interface {
	Bookmarks() ([]*bookmark.Bookmark, error)
	ChoiceSetByName(n string) (*choice_set.Set, error)
	ChoiceSets() ([]*choice_set.Set, error)
	CreateChoiceSet(
		name string,
		choices []string,
	) (*choice_set.Set, error)
	CreateSelectionField(
		name string,
		label string,
		choiceSet string,
	) (*custom_field.Field, error)
	CustomFields() ([]*custom_field.Field, error)
	DeleteChoiceSet(name string) error
	DeleteCustomField(name string) error
	SetChoiceSetChoices(
		name string,
		choices []string,
	) (*choice_set.Set, error)
	SetDeviceCustomFields(
		name string,
		fields map[string]any,
	) (*device.Device, error)
	SetVirtualMachineCustomFields(
		name string,
		fields map[string]any,
	) (*virtual_machine.Machine, error)
	ConfigurationContextByName(
		n string,
	) (*configuration_context.Context, error)
	ConfigurationContexts() ([]*configuration_context.Context, error)
	CreateConfigurationContext(
		name string,
		weight int32,
		payload map[string]any,
		a *configuration_assignment.Assignment,
	) (*configuration_context.Context, error)
	DeleteConfigurationContext(name string) error
	SetConfigurationContextAssignment(
		name string,
		a *configuration_assignment.Assignment,
	) (*configuration_context.Context, error)
	SetConfigurationContextPayload(
		name string,
		payload map[string]any,
	) (*configuration_context.Context, error)
	SetConfigurationContextWeight(
		name string,
		weight int32,
	) (*configuration_context.Context, error)
	ConfigurationTemplates() ([]*configuration_template.Template, error)
	CustomFieldByName(n string) (*custom_field.Field, error)
	CustomLinks() ([]*custom_link.Link, error)
	ExportTemplates() ([]*export_template.Template, error)
	MustBookmarks() []*bookmark.Bookmark
	MustChoiceSets() []*choice_set.Set
	MustConfigurationContexts() []*configuration_context.Context
	MustConfigurationTemplates() []*configuration_template.Template
	MustCustomFields() []*custom_field.Field
	MustCustomLinks() []*custom_link.Link
	MustExportTemplates() []*export_template.Template
	MustNotificationGroups() []*notification_group.Group
	NotificationGroups() ([]*notification_group.Group, error)
}
