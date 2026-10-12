package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/code"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/container"
	"github.com/funtimecoding/soil/pkg/github/image"
	"github.com/funtimecoding/soil/pkg/github/issue"
	"github.com/funtimecoding/soil/pkg/github/job"
	"github.com/funtimecoding/soil/pkg/github/release"
	"github.com/funtimecoding/soil/pkg/github/repository"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/github/tag"
	"github.com/funtimecoding/soil/pkg/github/workflow"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/google/go-github/v92/github"
	"testing"
	"time"
)

func TestConstant(t *testing.T) {
	assert.String(t, "X-GitHub-Event", constant.EventHeader)
	assert.String(t, "X-Hub-Signature-256", constant.SignatureHeader)
	assert.String(t, "closed", constant.Closed)
	assert.String(t, "open", constant.Open)
}

func TestStatusConstant(t *testing.T) {
	assert.String(t, "completed", constant.CompletedStatus)
	assert.String(t, "in_progress", constant.InProgressStatus)
	assert.String(t, "queued", constant.QueuedStatus)
	assert.String(t, "success", constant.SuccessConclusion)
	assert.String(t, "failure", constant.FailureConclusion)
	assert.String(t, "failed", constant.RunFailedConcern)
}

func TestWorkflowConstant(t *testing.T) {
	assert.String(t, "active", constant.WorkflowActiveState)
}

func TestCode(t *testing.T) {
	r := code.New(
		&github.CodeResult{
			SHA:  new(stringsConstant.UpperAlfa),
			Name: new(stringsConstant.UpperBravo),
			Path: new(stringsConstant.UpperCharlie),
		},
	)
	r.Raw = nil
	assert.Any(t, &code.Code{Hash: "Alfa", Name: "Bravo", Path: "Charlie"}, r)
}

func TestContainer(t *testing.T) {
	c := container.New(
		&github.Package{
			Name: new(stringsConstant.UpperAlfa),
			Repository: &github.Repository{
				Name: new(stringsConstant.UpperBravo),
			},
		},
	)
	c.Raw = nil
	assert.Any(t, &container.Container{Name: "Alfa", Repository: "Bravo"}, c)
}

func TestImage(t *testing.T) {
	meta, e := json.Marshal(
		github.PackageMetadata{
			PackageType: new(constant.ContainerPackageType),
			Container: &github.PackageContainerMetadata{
				Tags: []string{stringsConstant.UpperAlfa},
			},
		},
	)
	errors.PanicOnError(e)
	i := image.New(
		&github.PackageVersion{
			ID:        new(int64(1)),
			Name:      new(stringsConstant.UpperBravo),
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
			Title:   new(stringsConstant.UpperAlfa),
			HTMLURL: new(stringsConstant.UpperBravo),
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
			Name:      new(stringsConstant.UpperAlfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &job.Job{Name: "Alfa", CreatedAt: time.Time{}}, r)
}

func TestRelease(t *testing.T) {
	r := release.New(
		&github.RepositoryRelease{
			TagName:   stringsConstant.UpperAlfa,
			CreatedAt: github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &release.Release{Name: "Alfa", Create: time.Time{}}, r)
}

func TestRepository(t *testing.T) {
	r := repository.New(
		&github.Repository{
			Name:      new(stringsConstant.UpperAlfa),
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
					Name:   stringsConstant.UpperAlfa,
					Status: constant.CompletedStatus,
					Create: assert.NewDay(0),
				},
				{
					Name:   stringsConstant.UpperBravo,
					Status: constant.CompletedStatus,
					Create: assert.NewDay(1),
				},
				{
					Name:   stringsConstant.UpperCharlie,
					Status: constant.CompletedStatus,
					Create: assert.NewDay(2),
				},
			},
		).Name,
	)
}

func TestRun(t *testing.T) {
	r := run.New(
		&github.WorkflowRun{
			Name:       new(stringsConstant.UpperAlfa),
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
			Name:      new(stringsConstant.UpperAlfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &workflow.Workflow{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
