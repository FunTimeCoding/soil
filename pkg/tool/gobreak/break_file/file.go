package break_file

type File struct {
	Packages []string `yaml:"packages"`
	Tags     []string `yaml:"tags"`
	Tests    string   `yaml:"tests"`
	Breaks   []*Break `yaml:"breaks"`
}
