package option

type Certificate struct {
	Address            string
	ServiceTokens      []string
	PostgresLocator    string
	LitePath           string
	Project            string
	Branch             string
	AuthorityDirectory string
	SecretAuthority    string
	SecretPath         string
	Issuer             string
	ClientIdentifier   string
	ClientSecret       string
	EncryptionSecret   string
	PublicLocator      string
	Version            string
}
