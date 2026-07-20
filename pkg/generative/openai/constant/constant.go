package constant

const TokenEnvironment = "OPENAI_TOKEN"
const (
	NewSelector           = `a[data-testid="create-new-chat-button"]`
	ProfileSelector       = `[data-testid="accounts-profile-button"]`
	SettingsSelector      = `[data-testid="settings-menu-item"]`
	PersonalizeSelector   = `[data-testid="personalization-tab"]`
	MemoriesSelector      = `[class="btn relative btn-secondary btn-small"]`
	CloseMemoriesSelector = `div[role="dialog"] [data-testid="close-button"]`
	CloseSettingsSelector = `div[role="tablist"] [data-testid="close-button"]`
	PromptSelector        = `#prompt-textarea` // NewSelector not unique, requires index
	// CloseMemoriesSelector not unique, requires index
)

var UsefulAttributes = []string{"data-testid", "aria-label"}
