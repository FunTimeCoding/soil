package constant

import "regexp"

const (
	AnthropicTokenEnvironment = "ANTHROPIC_TOKEN"
	AnthropicVocabularyLink   = "https://raw.githubusercontent.com/rohangpta/ctoc/main/vocab.json"
)

const (
	AnthropicRoleUser      = "user"
	AnthropicRoleAssistant = "assistant"
)

const (
	ClaudeRecentMessageLimit = 50
	ClaudePendingCallLimit   = 64
)

const (
	ClaudeToolUseBlock    = "tool_use"
	ClaudeToolResultBlock = "tool_result"
)

var (
	ClaudeMarkupTagPattern = regexp.MustCompile(`<[^>]+>`)
	ClaudeAnsiPattern      = regexp.MustCompile(`\x1b\[[0-9;]*m`)
)

const AnthropicBodyElement = "body"
const (
	UsageMeterSession   = "Current session"
	UsageMeterAllModels = "All models"
	UsageMeterFable     = "Fable"
	UsageResetPrefix    = "Resets "
)
