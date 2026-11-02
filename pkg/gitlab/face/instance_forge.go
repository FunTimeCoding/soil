package face

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

type InstanceForge interface {
	AdminMode(
		x context.Context,
		on bool,
	) (*gitlab.Settings, error)
	CurrentUser(x context.Context) (*gitlab.User, error)
	FeatureDefinitions(
		x context.Context,
	) ([]*gitlab.FeatureDefinition, error)
	Features(x context.Context) ([]*gitlab.Feature, error)
	MustCurrentUser(x context.Context) *gitlab.User
	MustFeatureDefinitions(x context.Context) []*gitlab.FeatureDefinition
	MustFeatures(x context.Context) []*gitlab.Feature
	MustQuery(
		x context.Context,
		q string,
		response any,
	)
	Query(
		x context.Context,
		q string,
		response any,
	) error
	Setting(x context.Context) (*gitlab.Settings, error)
	Users(x context.Context) ([]*gitlab.User, error)
}
