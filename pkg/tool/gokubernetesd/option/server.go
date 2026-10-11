package option

type Server struct {
	Address         string
	ServiceTokens   []string
	ReadOnly        bool
	PostgresLocator string
	LitePath        string
}
