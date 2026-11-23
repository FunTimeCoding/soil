package trace_report

type Summary struct {
	States         int `json:"states"`
	Files          int `json:"files"`
	OrphanedStates int `json:"orphaned_states"`
	OrphanedFiles  int `json:"orphaned_files"`
}
