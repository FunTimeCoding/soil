package endpoint

type Endpoint struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Secure bool   `yaml:"secure"`
}
