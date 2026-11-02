package face

import "gitlab.com/gitlab-org/api/client-go/v3"

type InstanceForge interface {
	AdminMode(on bool) (*gitlab.Settings, error)
	CurrentUser() (*gitlab.User, error)
	FeatureDefinitions() ([]*gitlab.FeatureDefinition, error)
	Features() ([]*gitlab.Feature, error)
	MustCurrentUser() *gitlab.User
	MustFeatureDefinitions() []*gitlab.FeatureDefinition
	MustFeatures() []*gitlab.Feature
	MustQuery(
		q string,
		response any,
	)
	Query(
		q string,
		response any,
	) error
	Setting() (*gitlab.Settings, error)
	Users() ([]*gitlab.User, error)
}
