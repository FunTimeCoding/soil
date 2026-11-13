package break_file

type Break struct {
	Name   string  `yaml:"name"`
	Expect string  `yaml:"expect"`
	Edits  []*Edit `yaml:"edits"`
}
