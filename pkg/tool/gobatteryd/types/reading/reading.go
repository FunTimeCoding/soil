package reading

type Reading struct {
	Server    string
	Battery   string
	Attached  bool
	Variables map[string]string
}
