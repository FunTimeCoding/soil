package unit

import "testing"

func TestGuard(t *testing.T) {
	v := mountedReports(t, t.TempDir())
	v.VerifyBase(t)
	v.VerifyOpen(t, "/favicon.ico")
	v.VerifySignIn(t)
}
