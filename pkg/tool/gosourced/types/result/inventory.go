package result

type Inventory struct {
	Region string  `json:"region"`
	Total  int     `json:"total"`
	More   int     `json:"more"`
	Calls  []*Call `json:"calls"`
}
