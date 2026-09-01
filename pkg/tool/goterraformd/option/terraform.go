package option

type Terraform struct {
	Address         string
	MetricAddress   string
	Version         string
	Repository      string
	ClonePath       string
	TerraformPath   string
	PostgresLocator string
	LitePath        string
	StateNamespace  string
	StateLease      string
	Downstream      []string
}
