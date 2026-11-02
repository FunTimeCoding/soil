package response

type PodLog struct {
	Pod     string `json:"pod"`
	Content string `json:"content"`
}
