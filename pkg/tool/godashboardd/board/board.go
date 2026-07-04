package board

type Board struct {
	Connection Connection `yaml:"connection"`
	Top        []*Column  `yaml:"top"`
	Tail       Tail       `yaml:"tail"`
}
