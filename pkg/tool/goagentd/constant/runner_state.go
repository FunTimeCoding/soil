package constant

type RunnerState string

const (
	RunnerIdle    RunnerState = "idle"
	RunnerRunning RunnerState = "running"
	RunnerDone    RunnerState = "done"
)
