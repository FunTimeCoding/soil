package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/code"
	githubConstant "github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/container"
	"github.com/funtimecoding/soil/pkg/github/image"
	"github.com/funtimecoding/soil/pkg/github/issue"
	"github.com/funtimecoding/soil/pkg/github/job"
	"github.com/funtimecoding/soil/pkg/github/release"
	"github.com/funtimecoding/soil/pkg/github/repository"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/github/tag"
	"github.com/funtimecoding/soil/pkg/github/workflow"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/google/go-github/v92/github"
	"testing"
	"time"
)

func TestConstant(t *testing.T) {
	assert.String(t, "X-GitHub-Event", githubConstant.EventHeader)
	assert.String(t, "X-Hub-Signature-256", githubConstant.SignatureHeader)
	assert.String(t, "closed", githubConstant.Closed)
	assert.String(t, "open", githubConstant.Open)
}

func TestStatusConstant(t *testing.T) {
	assert.String(t, "completed", githubConstant.CompletedStatus)
	assert.String(t, "in_progress", githubConstant.InProgressStatus)
	assert.String(t, "queued", githubConstant.QueuedStatus)
	assert.String(t, "success", githubConstant.SuccessConclusion)
	assert.String(t, "failure", githubConstant.FailureConclusion)
	assert.String(t, "failed", githubConstant.RunFailedConcern)
}

func TestWorkflowConstant(t *testing.T) {
	assert.String(t, "active", githubConstant.WorkflowActiveState)
}

func TestCode(t *testing.T) {
	r := code.New(
		&github.CodeResult{
			SHA:  new(constant.UpperAlfa),
			Name: new(constant.UpperBravo),
			Path: new(constant.UpperCharlie),
		},
	)
	r.Raw = nil
	assert.Any(t, &code.Code{Hash: "Alfa", Name: "Bravo", Path: "Charlie"}, r)
}

func TestContainer(t *testing.T) {
	c := container.New(
		&github.Package{
			Name:       new(constant.UpperAlfa),
			Repository: &github.Repository{Name: new(constant.UpperBravo)},
		},
	)
	c.Raw = nil
	assert.Any(t, &container.Container{Name: "Alfa", Repository: "Bravo"}, c)
}

func TestImage(t *testing.T) {
	meta, e := json.Marshal(
		github.PackageMetadata{
			PackageType: new(githubConstant.ContainerPackageType),
			Container: &github.PackageContainerMetadata{
				Tags: []string{constant.UpperAlfa},
			},
		},
	)
	errors.PanicOnError(e)
	i := image.New(
		&github.PackageVersion{
			ID:        new(int64(1)),
			Name:      new(constant.UpperBravo),
			CreatedAt: &github.Timestamp{},
			Metadata:  meta,
		},
	)
	i.Raw = nil
	assert.Any(
		t,
		&image.Image{
			Identifier: 1,
			Digest:     "Bravo",
			Tags:       []string{"Alfa"},
			Create:     time.Time{},
		},
		i,
	)
}

func TestIssue(t *testing.T) {
	i := issue.New(
		&github.Issue{
			RepositoryURL: locator.New("api.github.com").Path(
				"/repos/funtimecoding/soil",
			).Pointer(),
			Title:   new(constant.UpperAlfa),
			HTMLURL: new(constant.UpperBravo),
		},
	)
	i.Raw = nil
	assert.Any(
		t,
		&issue.Issue{
			Repository: "funtimecoding/soil",
			Title:      "Alfa",
			Link:       "Bravo",
		},
		i,
	)
}

func TestJob(t *testing.T) {
	r := job.New(
		&github.WorkflowJob{
			Name:      new(constant.UpperAlfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &job.Job{Name: "Alfa", CreatedAt: time.Time{}}, r)
}

func TestRelease(t *testing.T) {
	r := release.New(
		&github.RepositoryRelease{
			TagName:   constant.UpperAlfa,
			CreatedAt: github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &release.Release{Name: "Alfa", Create: time.Time{}}, r)
}

func TestRepository(t *testing.T) {
	r := repository.New(
		&github.Repository{
			Name:      new(constant.UpperAlfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(
		t,
		&repository.Repository{Name: "Alfa", CreatedAt: time.Time{}},
		r,
	)
}

func TestRunLatest(t *testing.T) {
	assert.String(
		t,
		"Charlie",
		run.Latest(
			[]*run.Run{
				{
					Name:   constant.UpperAlfa,
					Status: githubConstant.CompletedStatus,
					Create: assert.NewDay(0),
				},
				{
					Name:   constant.UpperBravo,
					Status: githubConstant.CompletedStatus,
					Create: assert.NewDay(1),
				},
				{
					Name:   constant.UpperCharlie,
					Status: githubConstant.CompletedStatus,
					Create: assert.NewDay(2),
				},
			},
		).Name,
	)
}

func TestRun(t *testing.T) {
	r := run.New(
		&github.WorkflowRun{
			Name:       new(constant.UpperAlfa),
			CreatedAt:  &github.Timestamp{},
			Repository: &github.Repository{},
		},
	)
	r.Repository = nil
	r.Raw = nil
	assert.Any(
		t,
		&run.Run{
			MonitorIdentifier: "ghjob-0",
			Name:              "Alfa",
			Create:            time.Time{},
		},
		r,
	)
}

func TestTagLatest(t *testing.T) {
	assert.String(
		t,
		"v1.0.1",
		tag.Latest(
			[]*github.RepositoryTag{
				{Name: new("v1.0.0")},
				{Name: new("v1.0.1")},
			},
		).GetName(),
	)
}

func TestWorkflow(t *testing.T) {
	r := workflow.New(
		&github.Workflow{
			Name:      new(constant.UpperAlfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &workflow.Workflow{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
