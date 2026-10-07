package prometheus

type Prometheus struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Secure   bool   `yaml:"secure"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
