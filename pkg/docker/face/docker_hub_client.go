package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/docker/hub/tag"
)

type DockerHubClient interface {
	Close()
	MustTags(image string) []*tag.Tag
	Reachable(x context.Context) error
	Tags(image string) ([]*tag.Tag, error)
}
