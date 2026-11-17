package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/base"
	"testing"
)

func TestGuard(t *testing.T) {
	v := base.New(t).Served()
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/placements")
	v.VerifyGuarded(t, "/api/places")
	v.VerifyGuarded(t, "/api/sightings")
	v.VerifyFavicon(t)
	v.VerifySignIn(t)
}
