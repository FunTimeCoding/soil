package service

import "fmt"

func memoryPath(identifier int64) string {
	return fmt.Sprintf("memory/%d", identifier)
}
