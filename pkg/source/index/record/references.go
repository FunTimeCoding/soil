package record

type References struct {
	Targets map[string][]*Site `json:"targets"`
}
