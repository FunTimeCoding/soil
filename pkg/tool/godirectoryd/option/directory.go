package option

type Directory struct {
	Address          string
	ServiceTokens    []string
	Issuer           string
	ClientIdentifier string
	ClientSecret     string
	EncryptionSecret string
	PublicLocator    string
	Version          string
}
