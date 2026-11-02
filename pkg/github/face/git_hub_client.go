package face

import "github.com/funtimecoding/soil/pkg/face"

type GitHubClient interface {
	ActionClient
	PackageClient
	IssueClient
	AccountClient
	RepositoryClient
	face.Reachable
	Close()
}
