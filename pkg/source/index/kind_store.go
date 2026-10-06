package index

import "path"

func kindStore(
	base string,
	k *Kind,
) string {
	return path.Join(base, k.Name)
}
