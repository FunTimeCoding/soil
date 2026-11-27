package face

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/response"
	"github.com/funtimecoding/soil/pkg/face"
)

type LadybirdClient interface {
	face.Reachable
	Address() string
	Tree() ([]*response.Context, error)
	Navigate(context string, locator string) error
	Evaluate(context string, expression string) (json.RawMessage, error)
	CallFunction(
		context string,
		declaration string,
		arguments ...any,
	) (json.RawMessage, error)
	Type(context string, text string) error
	Screenshot(context string) ([]byte, error)
	Subscribe(events ...string) error
	Close()
}
