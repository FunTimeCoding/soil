package page_tree

type Tree struct {
	Title string  `json:"title"`
	Items []*Item `json:"tree"`
}
