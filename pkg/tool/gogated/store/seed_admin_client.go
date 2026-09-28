package store

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"golang.org/x/crypto/bcrypt"
)

func (s *Store) SeedAdminClient(
	identifier string,
	secret string,
	issuer string,
) error {
	if s.ClientExists(identifier) {
		return nil
	}

	hash, e := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)

	if e != nil {
		return e
	}

	return s.mapper.Create(
		client.New(
			identifier,
			string(hash),
			join.Empty(issuer, webConstant.CallbackPath),
			join.Space(
				constant.GrantAuthorizationCode,
				constant.GrantRefreshToken,
			),
			constant.ResponseCode,
			join.Space(constant.ScopeOpenIdentity, constant.ScopeOffline),
			false,
			constant.AuthMethodClientSecretPost,
		),
	).Error
}
