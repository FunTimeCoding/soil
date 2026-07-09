package repository

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
	"time"
)

func TestRepository(t *testing.T) {
	r := New(
		&github.Repository{
			Name:      new(upper.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &Repository{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
