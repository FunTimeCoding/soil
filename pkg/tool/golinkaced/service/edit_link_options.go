package service

type EditLinkOptions struct {
	Name        string
	Link        string
	Description string
	Tags        []string
	AddTags     []string
	RemoveTags  []string
	Lists       []string
	AddLists    []string
	RemoveLists []string
}
