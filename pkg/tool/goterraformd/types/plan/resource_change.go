package plan

type ResourceChange struct {
	Address string  `json:"address"`
	Type    string  `json:"type"`
	Change  *Change `json:"change"`
}
