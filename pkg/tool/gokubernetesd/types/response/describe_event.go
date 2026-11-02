package response

type DescribeEvent struct {
	Type    string `json:"type"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Age     string `json:"age"`
	Count   int32  `json:"count"`
}
