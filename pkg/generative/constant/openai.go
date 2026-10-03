package constant

const (
	OpenAITokenEnvironment = "OPENAI_TOKEN"
	OpenAIEmbedModel       = "text-embedding-nomic-embed-text-v1.5@f16"

	OpenAINewSelector           = `a[data-testid="create-new-chat-button"]`
	OpenAIProfileSelector       = `[data-testid="accounts-profile-button"]`
	OpenAISettingsSelector      = `[data-testid="settings-menu-item"]`
	OpenAIPersonalizeSelector   = `[data-testid="personalization-tab"]`
	OpenAIMemoriesSelector      = `[class="btn relative btn-secondary btn-small"]`
	OpenAICloseMemoriesSelector = `div[role="dialog"] [data-testid="close-button"]`
	OpenAICloseSettingsSelector = `div[role="tablist"] [data-testid="close-button"]`
	OpenAIPromptSelector        = `#prompt-textarea`

	OpenAINewIndex           int = 0
	OpenAICloseMemoriesIndex int = 2
)

var OpenAIUsefulAttributes = []string{"data-testid", "aria-label"}
