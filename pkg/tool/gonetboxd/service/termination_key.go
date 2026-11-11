package service

import "fmt"

func terminationKey(
	objectType string,
	identifier int64,
) string {
	return fmt.Sprintf("%s/%d", objectType, identifier)
}
