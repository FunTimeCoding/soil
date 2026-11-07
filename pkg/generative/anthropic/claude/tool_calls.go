package claude

import (
	"bufio"
	"encoding/json"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
)

func (c *Client) ToolCalls(sessionIdentifier string) []*tool_call.Call {
	path := filepath.Join(
		c.base,
		join.Empty(sessionIdentifier, soil.NotationLogExtension),
	)
	f, e := os.Open(path)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(f)
	var result []*tool_call.Call
	pending := map[string]*tool_call.Call{}
	scanner := bufio.NewScanner(f)
	scanner.Buffer(nil, soil.NotationScanBuffer)

	for scanner.Scan() {
		var line notation.Line

		if json.Unmarshal(scanner.Bytes(), &line) != nil ||
			line.Message == nil {
			continue
		}

		switch line.Type {
		case "assistant":
			for _, call := range lineCalls(&line) {
				pending[call.Identifier] = call
				result = append(result, call)
			}
		case "user":
			recordDecision(pending, &line)
		}
	}

	return result
}
