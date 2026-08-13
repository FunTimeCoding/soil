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

const ClaudeRecentMessageLimit = 50

var (
	ClaudeMarkupTagPattern = regexp.MustCompile(`<[^>]+>`)
	ClaudeAnsiPattern      = regexp.MustCompile(`\x1b\[[0-9;]*m`)
)

const AnthropicBodyElement = "body"

var (
	ResetPattern = regexp.MustCompile(`Resets?\s+(.+?)</span>`)
	ValuePattern = regexp.MustCompile(`aria-valuenow="(\d+)"`)
)
