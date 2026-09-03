package option

type Server struct {
	Address       string
	ServiceTokens []string
	ReadOnly      bool
	LitePath      string
	Version       string
}
