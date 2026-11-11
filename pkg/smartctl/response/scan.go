package response

type Scan struct {
	Devices []*Device `json:"devices"`
}
