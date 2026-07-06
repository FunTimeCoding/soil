package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

const Base = "/api/v4"
const (
	OpenedState  = "opened"
	PendingState = "pending"
	MergedState  = "merged"
	ClosedState  = "closed"
)

var RequestStates = []string{
	OpenedState,
	MergedState,
	ClosedState,
}

// Identifier Sort field
const Identifier = "id"

// Sort order
const (
	Ascending  = "asc"
	Descending = "desc"
)

// Job status
const (
	Success = "success"
	Failed  = "failed"
)

var DefaultListOptions = gitlab.ListOptions{Page: 0, PerPage: 100}

const (
	HostEnvironment    = "GITLAB_HOST"
	TokenEnvironment   = "GITLAB_TOKEN"
	GroupEnvironment   = "GITLAB_GROUP"
	ProjectEnvironment = "GITLAB_PROJECT"

	PrivateTokenHeader = "Private-Token"
	JobTokenHeader     = "Job-Token"

	LatestSuffix = ":latest"

	PerPage1000 int64 = 1000
	PerPage100  int64 = 100

	TokenHeader = "X-Gitlab-Token" // #nosec G101 not a hardcoded secret
)

// Environment variables during jobs
const (
	CommitTag         = "CI_COMMIT_TAG"
	InterfaceLocator  = "CI_API_V4_URL"
	JobToken          = "CI_JOB_TOKEN" // #nosec G101 not a hardcoded secret
	ProjectIdentifier = "CI_PROJECT_ID"
	ProjectNamespace  = "CI_PROJECT_NAMESPACE"
	QualifiedName     = "CI_SERVER_FQDN"
)

var (
	ErrorNotFound = errors.New("not found")
	Format        = option.Color.Copy()
	CheckFormat   = Format.Copy().Tag(tag.Project)
)
