package response

type Face struct {
	Identifier int   `json:"Id"`
	Colors     []int `json:"Colors"`
}
