package constant

import (
	"errors"
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
)

// Pull request state
const (
	All    = "all"
	Open   = "open"
	Closed = "closed"
)

var (
	ErrorNotFound = errors.New("not found")
	Format        = option.ExtendedColor.Copy()
)
