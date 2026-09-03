package option

type Credential struct {
	Address       string
	ServiceTokens []string
	Version       string
	Database      string
	Password      string
	RevealedField []string
}
