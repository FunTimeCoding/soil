package runner

import "gitlab.com/gitlab-org/api/client-go/v3"

type Runner struct {
	Identifier  int64
	Name        string
	Description string
	Status      string
	Type        string
	Shared      bool
	Online      bool
	Paused      bool
	Tags        []string
	Address     string // Loaded separately via GraphQL
	concern     []string
	RawList     *gitlab.Runner
	RawDetail   *gitlab.RunnerDetails
}
