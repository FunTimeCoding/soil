package scalar_or_list

import "go.yaml.in/yaml/v3"

func (s *Strings) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind == yaml.ScalarNode {
		*s = []string{value.Value}

		return nil
	}

	var list []string

	if e := value.Decode(&list); e != nil {
		return e
	}

	*s = list

	return nil
}
