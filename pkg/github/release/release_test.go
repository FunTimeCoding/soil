package release

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
	"time"
)

func TestRelease(t *testing.T) {
	r := New(
		&github.RepositoryRelease{
			TagName:   upper.Alfa,
			CreatedAt: github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(
		t,
		&Release{Name: "Alfa", Create: time.Time{}},
		r,
	)
}
