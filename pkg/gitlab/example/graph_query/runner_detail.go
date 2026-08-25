package graph_query

type runnerDetail struct {
	Identifier  string         `json:"id"`
	Description string         `json:"description"`
	Status      string         `json:"status"`
	RunnerType  string         `json:"runnerType"`
	Managers    runnerManagers `json:"managers"`
}
