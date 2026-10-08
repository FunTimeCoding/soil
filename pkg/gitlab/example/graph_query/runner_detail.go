package graph_query

type RunnerDetail struct {
	Identifier  string         `json:"id"`
	Description string         `json:"description"`
	Status      string         `json:"status"`
	RunnerType  string         `json:"runnerType"`
	Managers    RunnerManagers `json:"managers"`
}
