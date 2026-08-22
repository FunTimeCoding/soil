package service

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
)

func (s *Service) ResolveUser(name string) (*jira.User, error) {
	if name == constant.Unassign {
		return nil, nil
	}

	users, e := s.jira.FindUsers(name)

	if e != nil {
		return nil, e
	}

	if len(users) == 0 {
		return nil, not_found.Format("no user found matching: %s", name)
	}

	if len(users) == 1 {
		return &users[0], nil
	}

	for i := range users {
		if users[i].DisplayName == name {
			return &users[i], nil
		}
	}

	var names []string

	for _, u := range users {
		names = append(
			names,
			fmt.Sprintf("%s (%s)", u.DisplayName, u.AccountID),
		)
	}

	return nil, ambiguous.Format(
		"multiple users match '%s': %v - be more specific or use an account ID",
		name,
		names,
	)
}
