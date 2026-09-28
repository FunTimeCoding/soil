package signing_key

func New(
	identifier string,
	privateKey string,
	publicKey string,
	algorithm string,
) *SigningKey {
	return &SigningKey{
		Identifier: identifier,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Algorithm:  algorithm,
	}
}
