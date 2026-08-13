package metric

import "regexp"

func parse(s string) (string, map[string]string) {
	var name string
	nameMatches := regexp.MustCompile(`^([a-zA-Z_]\w*)`).FindStringSubmatch(s)

	if len(nameMatches) == 2 {
		name = nameMatches[1]
	}

	labels := make(map[string]string)
	labelMatches := regexp.MustCompile(`(\w+)="([^"]+)"`).FindAllStringSubmatch(
		s,
		-1,
	)

	for _, match := range labelMatches {
		if len(match) == 3 {
			label := match[1]
			value := match[2]
			labels[label] = value
		}
	}

	return name, labels
}
