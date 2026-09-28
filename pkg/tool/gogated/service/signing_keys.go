package service

func (s *Service) SigningKeys() *SigningKeySet {
	return &SigningKeySet{
		Keys: []*SigningKey{
			publicKeyToSigningKey(&s.signingKey.PublicKey, s.keyIdentifier),
		},
	}
}
