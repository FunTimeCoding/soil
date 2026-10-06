package index

import "github.com/funtimecoding/soil/pkg/source/constant"

func (w *Workspace) read(
	path string,
	key string,
) bool {
	v, okay := w.fetch(
		constant.IndexPackageKind,
		path,
		key,
		NewPackageRecord(""),
	)

	if !okay {
		return false
	}

	w.fingerprints[path] = v.(*PackageRecord).Fingerprint
	complete := true

	for _, k := range w.kinds {
		if _, known := w.facts[k.Name][path]; known {
			continue
		}

		value, found := w.fetch(
			kindStore(constant.IndexFactsKind, k),
			path,
			kindKey(key, k),
			k.Value(),
		)

		if found {
			w.facts[k.Name][path] = value

			continue
		}

		complete = false
	}

	return complete
}
