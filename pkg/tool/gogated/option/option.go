package option

type Option struct {
	Address               string
	PostgresLocator       string
	LitePath              string
	Secret                string
	SuperUserMail         string
	SuperUserPassword     string
	Issuer                string
	AdminClientIdentifier string
	AdminClientSecret     string
	ServiceTokens         []string
	Version               string
}
