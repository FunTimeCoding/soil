package receipt

type Receipt struct {
	Identifier    string
	Name          string
	Events        int64
	EventMetadata int64
	Completions   int64
	Summaries     int64
	Labels        int64
	Pulses        int64
	ContextLoads  int64
	TrackerStates int64
	Queue         int64
	Notifications int64
	Transcript    string
	Sources       []string
}
