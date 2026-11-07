package constant

const ModernDiscoverRequest = `{"jsonrpc":"2.0","id":1,"method":"server/discover","params":{"_meta":{"io.modelcontextprotocol/protocolVersion":"2026-07-28","io.modelcontextprotocol/clientCapabilities":{}}}}`
const (
	FixtureScratchDirectory  = "tmp"
	FixtureDocumentDirectory = "doc/guide"
)

var FixtureAllowlist = []string{
	"Bash(app run:*)",
	"Bash(awk:*)",
	"Bash(cat:*)",
	"Bash(cd:*)",
	"Bash(continue)",
	"Bash(echo:*)",
	"Bash(find:*)",
	"Bash(git diff:*)",
	"Bash(grep:*)",
	"Bash(gsed:*)",
	"Bash(head:*)",
	"Bash(ls:*)",
	"Bash(read:*)",
	"Bash(sed:*)",
	"Bash(set:*)",
	"Bash(wc:*)",
	"Read",
}
