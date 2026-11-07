package unit

import "testing"

func TestGuard(t *testing.T) {
	v := mountedReports(t, t.TempDir())
	v.VerifyBase(t)
	v.VerifyFavicon(t)
	v.VerifySignIn(t)
}
