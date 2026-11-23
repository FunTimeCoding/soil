package graph_tester

import "fmt"

func DeviceCode(expiry int) string {
	return fmt.Sprintf(
		`{"device_code":"device-1","user_code":"USER-1","verification_uri":"https://microsoft.com/devicelogin","expires_in":%d,"interval":1}`,
		expiry,
	)
}
