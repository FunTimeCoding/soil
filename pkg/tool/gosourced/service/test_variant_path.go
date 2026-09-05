package service

import "fmt"

func testVariantPath(packagePath string) string {
	return fmt.Sprintf("%s_test", packagePath)
}
