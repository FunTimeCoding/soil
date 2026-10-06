package layout

type Tail struct {
	Columns  int        `yaml:"columns"`
	Sections []*Section `yaml:"sections"`
}
