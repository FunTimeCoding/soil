package repository

import (
	"github.com/google/go-github/v89/github"
	"time"
)

func New(v *github.Repository) *Repository {
	var create time.Time

	if v.CreatedAt != nil {
		create = v.CreatedAt.Time
	}

	return &Repository{
		Identifier: v.GetID(),
		Owner:      v.GetOwner().GetLogin(),
		Name:       v.GetName(),
		FullName:   v.GetFullName(),
		CreatedAt:  create,
		Raw:        v,
	}
}
