package face

import (
	"context"
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
	Branches(
		x context.Context,
		project int64,
	) ([]*branch.Branch, error)
	MustBranches(
		x context.Context,
		project int64,
	) []*branch.Branch
	Tags(
		x context.Context,
		project int64,
	) ([]*tag.Tag, error)
	MustTags(
		x context.Context,
		project int64,
	) []*tag.Tag
	CreateTag(
		x context.Context,
		project int64,
		name string,
		reference string,
		message string,
	) (*tag.Tag, error)
	Commit(
		x context.Context,
		project int64,
		branch string,
		text string,
		path string,
		content string,
		update bool,
	) (*commit.Commit, error)
	MustCommit(
		x context.Context,
		project int64,
		branch string,
		text string,
		path string,
		content string,
		update bool,
	) *commit.Commit
	ReadCommit(
		x context.Context,
		project int64,
		sha string,
	) (*commit.Commit, error)
	ListCommits(
		x context.Context,
		project int64,
		reference string,
		limit int64,
	) ([]*commit.Commit, error)
	CommitDiff(
		x context.Context,
		project int64,
		sha string,
		limit int64,
	) ([]*diff.Diff, error)
	CompareFiles(
		x context.Context,
		project int64,
		from string,
		to string,
	) ([]string, error)
	File(
		x context.Context,
		project int64,
		branch string,
		name string,
	) (*file.File, error)
	MustFile(
		x context.Context,
		project int64,
		branch string,
		name string,
	) *file.File
	Tree(
		x context.Context,
		project int64,
		path string,
		reference string,
		recursive bool,
		limit int64,
	) ([]*tree.Node, error)
	CreateBranch(
		x context.Context,
		project int64,
		name string,
		reference string,
	) (*branch.Branch, error)
	CommitActions(
		x context.Context,
		project int64,
		branch string,
		message string,
		v []*gitlab.CommitActionOptions,
	) (*commit.Commit, error)
	MustCommitActions(
		x context.Context,
		project int64,
		branch string,
		message string,
		v []*gitlab.CommitActionOptions,
	) *commit.Commit
	Compare(
		x context.Context,
		project int64,
		from, to string,
	) (*gitlab.Compare, error)
	DeleteTag(
		x context.Context,
		project int64,
		name string,
	) error
	FileExists(
		x context.Context,
		p *project.Project,
		branch string,
		file string,
	) (bool, error)
	MustDeleteTag(
		x context.Context,
		project int64,
		name string,
	)
	MustFileExists(
		x context.Context,
		p *project.Project,
		branch string,
		file string,
	) bool
}
