package chain

type Chain struct {
	Environment string   `json:"environment"`
	Target      string   `json:"target"`
	Live        bool     `json:"live"`
	States      []string `json:"states"`
}
