package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/release"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
	"time"
)

func TestRelease(t *testing.T) {
	r := release.New(
		&github.RepositoryRelease{
			TagName:   upper.Alfa,
			CreatedAt: github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(
		t,
		&release.Release{Name: "Alfa", Create: time.Time{}},
		r,
	)
}
