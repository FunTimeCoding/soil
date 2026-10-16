package host_mapping

func (x *Index) existing(address string) []string {
	if _, okay := x.resources[address]; !okay {
		return nil
	}

	return []string{address}
}
