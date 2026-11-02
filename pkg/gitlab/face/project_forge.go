package face

import (
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

type ProjectForge interface {
	Project(identifier int64) (*project.Project, error)
	Projects() ([]*project.Project, error)
	MustProjects() []*project.Project
	SearchProject(query string) ([]*project.Project, error)
	ResolveProject(identifier string) (int64, error)
	GroupByName(s string) (*gitlab.Group, error)
	GroupProjects(identifier ...int64) ([]*project.Project, error)
	Groups() ([]*gitlab.Group, error)
	MustProject(identifier int64) *project.Project
	MustProjectByName(
		namespace string,
		name string,
	) *project.Project
	MustSearchBlob(query string) []*gitlab.Blob
	MustSearchProject(query string) []*project.Project
	PipelineProjects() ([]*project.Project, error)
	ProjectByName(
		namespace string,
		name string,
	) (*project.Project, error)
	ProjectsWithFile(
		path string,
		caseInsensitive bool,
	) ([]*project.Project, error)
	SearchBlob(query string) ([]*gitlab.Blob, error)
}
