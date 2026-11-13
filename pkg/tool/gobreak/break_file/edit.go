package break_file

type Edit struct {
	File    string `yaml:"file"`
	Find    string `yaml:"find"`
	Replace string `yaml:"replace"`
}
