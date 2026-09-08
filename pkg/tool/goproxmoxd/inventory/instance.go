package inventory

type Instance struct {
	Name        string `yaml:"name"`
	Index       int    `yaml:"index"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Token       string `yaml:"token"`
	Secret      string `yaml:"secret"`
	Untrusted   bool   `yaml:"untrusted"`
	Timeout     string `yaml:"timeout"`
	SSHUser     string `yaml:"ssh_user"`
	SSHPassword string `yaml:"ssh_password"`
}
