package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/reservation"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestReservation(t *testing.T) {
	assert.NotNil(t, reservation.New(&netbox.RackReservation{}))
}
