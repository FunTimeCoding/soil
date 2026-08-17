package rule

type Rule struct {
	Identifier      string
	Enabled         bool
	Sequence        string
	Interface       string
	Direction       string
	Action          string
	Protocol        string
	SourceNet       string
	SourcePort      string
	DestinationNet  string
	DestinationPort string
	Log             bool
	Automatic       bool
	Description     string
	Categories      string
}
