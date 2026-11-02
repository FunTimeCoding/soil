package face

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"github.com/funtimecoding/soil/pkg/gitlab/diff"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"github.com/funtimecoding/soil/pkg/gitlab/tree"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

type RepositoryForge interface {
	Branches(project int64) ([]*branch.Branch, error)
	MustBranches(project int64) []*branch.Branch
	Tags(project int64) ([]*tag.Tag, error)
	MustTags(project int64) []*tag.Tag
	CreateTag(
		project int64,
		name string,
		reference string,
		message string,
	) (*tag.Tag, error)
	Commit(
		project int64,
		branch string,
		text string,
		path string,
		content string,
		update bool,
	) (*commit.Commit, error)
	MustCommit(
		project int64,
		branch string,
		text string,
		path string,
		content string,
		update bool,
	) *commit.Commit
	ReadCommit(
		project int64,
		sha string,
	) (*commit.Commit, error)
	ListCommits(
		project int64,
		reference string,
		limit int64,
	) ([]*commit.Commit, error)
	CommitDiff(
		project int64,
		sha string,
		limit int64,
	) ([]*diff.Diff, error)
	CompareFiles(
		project int64,
		from string,
		to string,
	) ([]string, error)
	File(
		project int64,
		branch string,
		name string,
	) (*file.File, error)
	MustFile(
		project int64,
		branch string,
		name string,
	) *file.File
	Tree(
		project int64,
		path string,
		reference string,
		recursive bool,
		limit int64,
	) ([]*tree.Node, error)
	CreateBranch(
		project int64,
		name string,
		reference string,
	) (*branch.Branch, error)
	CommitActions(
		project int64,
		branch string,
		message string,
		v []*gitlab.CommitActionOptions,
	) (*commit.Commit, error)
	MustCommitActions(
		project int64,
		branch string,
		message string,
		v []*gitlab.CommitActionOptions,
	) *commit.Commit
	Compare(
		project int64,
		from, to string,
	) (*gitlab.Compare, error)
	DeleteTag(
		project int64,
		name string,
	) error
	FileExists(
		p *project.Project,
		branch string,
		file string,
	) (bool, error)
	MustDeleteTag(
		project int64,
		name string,
	)
	MustFileExists(
		p *project.Project,
		branch string,
		file string,
	) bool
}
