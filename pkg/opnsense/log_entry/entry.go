package log_entry

type Entry struct {
	Timestamp       string
	Interface       string
	Action          string
	Direction       string
	Protocol        string
	Source          string
	SourcePort      string
	Destination     string
	DestinationPort string
	RuleNumber      string
	RuleIdentifier  string
	Label           string
}
