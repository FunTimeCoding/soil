package page_tree

type Item struct {
	Index    *int    `json:"i"`
	Role     string  `json:"r"`
	Name     string  `json:"n"`
	Value    string  `json:"v"`
	Children []*Item `json:"c"`
}
