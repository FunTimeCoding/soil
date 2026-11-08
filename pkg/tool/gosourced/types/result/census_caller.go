package result

type CensusCaller struct {
	Kind     string `json:"kind"`
	Package  string `json:"package"`
	Function string `json:"function"`
	Count    int    `json:"count"`
}
