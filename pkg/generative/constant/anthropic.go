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
const (
	ClaudeDirectory              = ".claude"
	ClaudeCredentialFile         = ".credentials.json"
	AnthropicCredentialService   = "Claude Code-credentials"
	AnthropicUtilizationLink     = "https://api.anthropic.com/api/oauth/usage"
	AnthropicUtilizationBeta     = "oauth-2025-04-20"
	AnthropicBetaHeader          = "anthropic-beta"
	AnthropicBaseLinkEnvironment = "ANTHROPIC_BASE_URL"
)
const (
	AnthropicLimitWeekly       = "weekly_all"
	AnthropicLimitWeeklyScoped = "weekly_scoped"
)
