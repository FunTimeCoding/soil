package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func (s *Service) RegisterClient(
	r *RegisterClientRequest,
) (*RegisterClientResponse, error) {
	clientIdentifier := uuid.New().String()
	secret := uuid.New().String()
	hash, e := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)

	if e != nil {
		return nil, e
	}

	authMethod := r.TokenEndpointAuthMethod

	if authMethod == "" {
		authMethod = constant.AuthMethodClientSecretBasic
	}

	e = s.store.CreateClient(
		client.New(
			clientIdentifier,
			string(hash),
			strings.Join(r.RedirectLocators, "\n"),
			strings.Join(r.GrantTypes, " "),
			strings.Join(r.ResponseTypes, " "),
			strings.Join(r.Scopes, " "),
			authMethod == constant.AuthMethodNone,
			authMethod,
		),
	)

	if e != nil {
		return nil, e
	}

	return &RegisterClientResponse{
		ClientIdentifier: clientIdentifier,
		ClientSecret:     secret,
		RedirectLocators: r.RedirectLocators,
	}, nil
}
