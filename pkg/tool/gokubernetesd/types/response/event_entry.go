package response

type EventEntry struct {
	Namespace string `json:"namespace"`
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	Object    string `json:"object"`
	Message   string `json:"message"`
	Count     int32  `json:"count"`
	Age       string `json:"age"`
}
