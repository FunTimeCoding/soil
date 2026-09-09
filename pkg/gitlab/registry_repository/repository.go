package registry_repository

import (
	"gitlab.com/gitlab-org/api/client-go/v2"
	"time"
)

type Repository struct {
	Identifier        int64
	ProjectIdentifier int64
	Name              string
	Path              string
	Location          string
	Create            *time.Time
	TagCount          int64
	Raw               *gitlab.RegistryRepository
}
