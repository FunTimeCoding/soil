package service

type SigningKey struct {
	KeyType       string `json:"kty"`
	Use           string `json:"use"`
	KeyIdentifier string `json:"kid"`
	Algorithm     string `json:"alg"`
	N             string `json:"n"`
	E             string `json:"e"`
}
