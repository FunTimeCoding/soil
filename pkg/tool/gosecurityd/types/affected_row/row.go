package affected_row

type Row struct {
	System    string
	Coverage  string
	Priority  int
	Labels    map[string]string
	Source    string
	Installed string
	Fixed     string
	Candidate string
	Kind      string
}
