package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set_record"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set_request"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net/http"
)

func (c *Client) SetChoiceSetChoices(
	name string,
	choices []string,
) (*choice_set.Set, error) {
	v, e := c.ChoiceSetByName(name)

	if e != nil {
		return nil, e
	}

	result := choice_set_record.New()

	if f := c.rawRequest(
		http.MethodPatch,
		fmt.Sprintf("%s%d/", constant.ChoiceSetPath, v.Identifier),
		choice_set_request.New(choices),
		result,
	); f != nil {
		return nil, f
	}

	return choice_set.New(result), nil
}
