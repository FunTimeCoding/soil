package response

func NewSigningKeySet(keys []*SigningKey) *SigningKeySet {
	return &SigningKeySet{Keys: keys}
}
