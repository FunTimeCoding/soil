package response

func NewSigningKey(
	keyType string,
	use string,
	keyIdentifier string,
	algorithm string,
	n string,
	e string,
) *SigningKey {
	return &SigningKey{
		KeyType:       keyType,
		Use:           use,
		KeyIdentifier: keyIdentifier,
		Algorithm:     algorithm,
		N:             n,
		E:             e,
	}
}
