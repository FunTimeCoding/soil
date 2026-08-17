package source_nat

type Rule struct {
	Identifier     string
	Enabled        bool
	Interface      string
	Protocol       string
	SourceNet      string
	DestinationNet string
	Target         string
	TargetPort     string
	Log            bool
	Description    string
}
