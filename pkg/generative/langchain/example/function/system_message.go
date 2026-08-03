package function

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func systemMessage() string {
	b, e := json.Marshal(constant.LangchainExampleFunctions)
	errors.PanicOnError(e)

	return fmt.Sprintf(
		`You have access to the following tools:

%s

To use a tool, respond with a JSON object with the following structure:
{
	"tool": <name of the called tool>,
	"tool_input": <parameters for the tool matching the above JSON schema>
}
`,
		string(b),
	)
}
