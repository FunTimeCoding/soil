package connection

type Connection struct {
	Host     string `yaml:"host"`
	Secure   bool   `yaml:"secure"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
