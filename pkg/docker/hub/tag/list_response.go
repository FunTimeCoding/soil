package tag

type ListResponse struct {
	Count   int        `json:"count"`
	Results []Response `json:"results"`
}
