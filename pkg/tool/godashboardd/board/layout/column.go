package layout

type Column struct {
	Name     string     `yaml:"name"`
	Sections []*Section `yaml:"sections"`
}
