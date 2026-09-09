package image

import (
	"gitlab.com/gitlab-org/api/client-go/v2"
	"time"
)

type Image struct {
	Name     string
	Path     string
	Location string
	Digest   string
	Revision string
	Create   *time.Time
	Size     int64
	Raw      *gitlab.RegistryRepositoryTag
}
