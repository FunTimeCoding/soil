package argument

type EditLink struct {
	Identifier  int      `json:"identifier"`
	Name        string   `json:"name"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	AddTags     []string `json:"add_tags"`
	RemoveTags  []string `json:"remove_tags"`
	Lists       []string `json:"lists"`
	AddLists    []string `json:"add_lists"`
	RemoveLists []string `json:"remove_lists"`
}
