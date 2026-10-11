package image

import "regexp"

func FindNames(input string) []string {
	var result []string

	for _, match := range regexp.MustCompile(
		`debian-(\d+\.\d+\.\d+)-arm64-netinst\.iso`,
	).FindAllStringSubmatch(input, -1) {
		result = append(result, match[0])
	}

	return result
}
