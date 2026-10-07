package response

type Emblem struct {
	Background Face  `json:"Background"`
	Foreground Face  `json:"Foreground"`
	Flags      []any `json:"Flags"`
}
