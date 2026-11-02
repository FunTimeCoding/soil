package face

import "context"

type GitLabClient interface {
	ProjectForge
	RepositoryForge
	MergeRequestForge
	PipelineForge
	JobForge
	RunnerForge
	RegistryForge
	VariableForge
	IssueForge
	InstanceForge
	Close()
	Reachable(x context.Context) error
}
