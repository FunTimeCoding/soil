package break_result

type Result struct {
	Name   string   `json:"name"`
	Status string   `json:"status"`
	Tests  []string `json:"tests,omitempty"`
	Detail string   `json:"detail,omitempty"`
	Expect string   `json:"expect,omitempty"`
}
