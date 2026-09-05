package unit

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
	"testing"
)

func TestHostWrapperMarshal(t *testing.T) {
	h := request.New()
	h.Host = "example"
	h.Domain = "test"
	h.Address = "10.0.0.1"
	h.HardwareAddress = "02:00:00:00:00:01"
	b, e := json.Marshal(request.NewHostWrapper(h))
	assert.Nil(t, e)
	assert.String(
		t,
		`{"host":{"host":"example","domain":"test","ip":"10.0.0.1","hwaddr":"02:00:00:00:00:01"}}`,
		string(b),
	)
}

func TestHostRequestOmitsUnsetFields(t *testing.T) {
	h := request.New()
	h.Description = "reservation"
	b, e := json.Marshal(request.NewHostWrapper(h))
	assert.Nil(t, e)
	assert.String(t, `{"host":{"descr":"reservation"}}`, string(b))
}
