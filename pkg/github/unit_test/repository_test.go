package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/repository"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
	"time"
)

func TestRepository(t *testing.T) {
	r := repository.New(
		&github.Repository{
			Name:      new(upper.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &repository.Repository{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
