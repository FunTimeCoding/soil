package response

type Report struct {
	Warnings []string `json:"Warnings"`
	Added    []string `json:"Added"`
	Removed  []string `json:"Removed"`
}
