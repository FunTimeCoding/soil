package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field_record"
	"net/http"
)

func (c *Client) CustomFields() ([]*custom_field.Field, error) {
	var result custom_field_record.Page

	if e := c.rawRequest(
		http.MethodGet,
		constant.CustomFieldListPath,
		nil,
		&result,
	); e != nil {
		return nil, e
	}

	return custom_field.NewSlice(result.Results), nil
}
