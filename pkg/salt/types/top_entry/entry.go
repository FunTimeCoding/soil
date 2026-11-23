package top_entry

type Entry struct {
	Environment string `json:"environment"`
	Target      string `json:"target"`
	State       string `json:"state"`
	Live        bool   `json:"live"`
}
