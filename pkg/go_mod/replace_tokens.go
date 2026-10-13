package go_mod

import "golang.org/x/mod/modfile"

func replaceTokens(r *modfile.Replace) []string {
	result := []string{modfile.AutoQuote(r.Old.Path)}

	if r.Old.Version != "" {
		result = append(result, r.Old.Version)
	}

	result = append(result, "=>", modfile.AutoQuote(r.New.Path))

	if r.New.Version != "" {
		result = append(result, r.New.Version)
	}

	return result
}
