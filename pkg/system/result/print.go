package result

import "fmt"

func (r *Result) Print() {
	if r.Exit != 0 || r.Error != nil {
		fmt.Printf("Error (%d): %v\n", r.Exit, r.Error)
	}

	if r.OutputString != "" {
		fmt.Printf("Stdout:\n%s\n", r.OutputString)
	}

	if r.ErrorString != "" {
		fmt.Printf("Stderr:\n%s\n", r.ErrorString)
	}
}
