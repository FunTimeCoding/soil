package response

type Meta struct {
	TemplateCredsSetupCompleted bool   `json:"templateCredsSetupCompleted,omitzero"`
	TemplateIdentifier          string `json:"templateId,omitempty"`
}
