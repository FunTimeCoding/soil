package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

type CustomizationClient interface {
	CreateChoiceSet(
		x context.Context,
		name string,
		choices []string,
	) (client.ChoiceSet, error)
	CreateConfigurationContext(
		x context.Context,
		body client.CreateConfigurationContextRequest,
	) (client.ConfigurationContext, error)
	CreateCustomField(
		x context.Context,
		name string,
		label string,
		choiceSet string,
	) (client.CustomField, error)
	DeleteChoiceSet(
		x context.Context,
		name string,
	) error
	DeleteConfigurationContext(
		x context.Context,
		name string,
	) error
	DeleteCustomField(
		x context.Context,
		name string,
	) error
	GetChoiceSet(
		x context.Context,
		name string,
	) (client.ChoiceSet, error)
	GetConfigurationContext(
		x context.Context,
		name string,
	) (client.ConfigurationContext, error)
	ListChoiceSets(
		x context.Context,
	) ([]*client.ChoiceSet, error)
	ListConfigurationContexts(
		x context.Context,
	) ([]*client.ConfigurationContext, error)
	ListCustomFields(
		x context.Context,
	) ([]*client.CustomField, error)
	UpdateChoiceSet(
		x context.Context,
		name string,
		choices []string,
	) (client.ChoiceSet, error)
	UpdateConfigurationContext(
		x context.Context,
		name string,
		body client.UpdateConfigurationContextRequest,
	) (client.ConfigurationContext, error)
}
