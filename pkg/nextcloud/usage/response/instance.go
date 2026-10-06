package response

type Instance struct {
	Storage Storage `json:"storage"`
	Shares  Shares  `json:"shares"`
}
