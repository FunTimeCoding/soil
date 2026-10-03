package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

const (
	TokenEnvironment     = "GITHUB_TOKEN" // #nosec G101 not a hardcoded secret
	RunEnvironment       = "GITHUB_RUN_ID"
	ReferenceEnvironment = "GITHUB_REF_NAME"

	DelveNamespace  = "go-delve"
	DelveRepository = "delve"

	Namespace  = "funtimecoding"
	Repository = "soil"

	EventHeader     = "X-GitHub-Event"
	SignatureHeader = "X-Hub-Signature-256"

	ContainerPackageType = "container"

	MaximumPerPage = 100
)

const (
	All    = "all"
	Open   = "open"
	Closed = "closed"
)

var (
	Format         = constant.ExtendedColorFormat.Copy()
	NotationFormat = option.New()
)

const (
	CompletedStatus  = "completed"
	QueuedStatus     = "queued"
	InProgressStatus = "in_progress"
)

const (
	SuccessConclusion = "success"
	FailureConclusion = "failure"
)

const RunFailedConcern = "failed"
const WorkflowActiveState = "active"
const NoTags = "no tags"
