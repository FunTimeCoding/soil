package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"strings"
)

func (s *Service) UpdateClient(
	identifier string,
	redirectLocators []string,
	scopes []string,
	grantTypes []string,
	responseTypes []string,
	tokenEndpointAuthMethod string,
) (*client.Client, error) {
	row, e := s.store.ClientRow(identifier)

	if e != nil {
		return nil, e
	}

	if len(redirectLocators) > 0 {
		row.RedirectLocators = strings.Join(redirectLocators, "\n")
	}

	if len(scopes) > 0 {
		row.Scopes = strings.Join(scopes, " ")
	}

	if len(grantTypes) > 0 {
		row.GrantTypes = strings.Join(grantTypes, " ")
	}

	if len(responseTypes) > 0 {
		row.ResponseTypes = strings.Join(responseTypes, " ")
	}

	if tokenEndpointAuthMethod != "" {
		row.TokenEndpointAuthMethod = tokenEndpointAuthMethod
	}

	if e = s.store.UpdateClient(row); e != nil {
		return nil, e
	}

	return row, nil
}
