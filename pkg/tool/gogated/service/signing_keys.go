package service

import "github.com/funtimecoding/soil/pkg/tool/gogated/types/response"

func (s *Service) SigningKeys() *response.SigningKeySet {
	return response.NewSigningKeySet(
		[]*response.SigningKey{
			publicKeyToSigningKey(&s.signingKey.PublicKey, s.keyIdentifier),
		},
	)
}
