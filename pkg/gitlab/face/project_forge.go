package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

type ProjectForge interface {
	Project(
		x context.Context,
		identifier int64,
	) (*project.Project, error)
	Projects(x context.Context) ([]*project.Project, error)
	MustProjects(x context.Context) []*project.Project
	SearchProject(
		x context.Context,
		query string,
	) ([]*project.Project, error)
	ResolveProject(
		x context.Context,
		identifier string,
	) (int64, error)
	GroupByName(
		x context.Context,
		s string,
	) (*gitlab.Group, error)
	GroupProjects(
		x context.Context,
		identifier ...int64,
	) ([]*project.Project, error)
	Groups(x context.Context) ([]*gitlab.Group, error)
	MustProject(
		x context.Context,
		identifier int64,
	) *project.Project
	MustProjectByName(
		x context.Context,
		namespace string,
		name string,
	) *project.Project
	MustSearchBlob(
		x context.Context,
		query string,
	) []*gitlab.Blob
	MustSearchProject(
		x context.Context,
		query string,
	) []*project.Project
	PipelineProjects(x context.Context) ([]*project.Project, error)
	ProjectByName(
		x context.Context,
		namespace string,
		name string,
	) (*project.Project, error)
	ProjectsWithFile(
		x context.Context,
		path string,
		caseInsensitive bool,
	) ([]*project.Project, error)
	SearchBlob(
		x context.Context,
		query string,
	) ([]*gitlab.Blob, error)
}
