package kernel_offer

type Offer struct {
	Version  string   `json:"version"`
	Priority int      `json:"priority"`
	Origins  []string `json:"origins"`
}
