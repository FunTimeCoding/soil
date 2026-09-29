package argument

type ListLinks struct {
	List   int `json:"list"`
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
