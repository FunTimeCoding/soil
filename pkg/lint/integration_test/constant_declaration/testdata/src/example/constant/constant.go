package constant

type Severity string

const (
	SeverityLow  Severity = "low"
	SeverityHigh Severity = "high"
)

type Record struct {
	Name string
}

var Records = []Record{
	{Name: "flagged type above, allowed table here"},
}

func Helper() bool {
	return true
}
