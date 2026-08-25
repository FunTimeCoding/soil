package checklist_item

type Item struct {
	Identifier string `json:"id"`
	Text       string `json:"text"`
	Completed  bool   `json:"completed"`
}
