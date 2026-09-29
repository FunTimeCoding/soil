package argument

type AddLink struct {
	Name string   `json:"name"`
	Link string   `json:"link"`
	List string   `json:"list"`
	Tags []string `json:"tags"`
}
