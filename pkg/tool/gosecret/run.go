package gosecret

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
)

func Run(
	directory string,
	mode constant.Mode,
) error {
	secrets, e := FindSecrets(directory)

	if e != nil {
		return fmt.Errorf("find secrets: %w", e)
	}

	if len(secrets) == 0 {
		fmt.Println("No secrets found")

		return nil
	}

	fmt.Printf("Found %d secret manifest(s)\n", len(secrets))
	var (
		processed  int
		inSync     int
		fails      []string
		mismatches []string
	)

	for _, path := range secrets {
		var s *SecretResult
		var f error

		if mode == constant.Encode {
			s, f = EncodeSecret(path)
		} else {
			s, f = ProcessSecret(path, mode == constant.Check)
		}

		if f != nil {
			fails = append(fails, fmt.Sprintf("%s: %v", path, f))

			continue
		}

		if s == nil {
			continue
		}

		processed++

		switch mode {
		case constant.Check:
			if !s.InSync {
				mismatches = append(
					mismatches,
					fmt.Sprintf("%s -> %s", path, s.DecodedPath),
				)
			}
		case constant.Encode:
			if s.InSync {
				inSync++
			} else {
				fmt.Printf("✓ %s <- %s\n", path, s.DecodedPath)
			}
		default:
			fmt.Printf("✓ %s -> %s\n", path, s.DecodedPath)
		}
	}

	if len(fails) > 0 {
		errors.Printf("\nErrors:\n")

		for _, m := range fails {
			errors.Printf("  ✗ %s\n", m)
		}
	}

	switch mode {
	case constant.Check:
		if len(mismatches) > 0 {
			errors.Printf("\nMismatched secrets (%d):\n", len(mismatches))

			for _, m := range mismatches {
				errors.Printf("  ✗ %s\n", m)
			}

			return fmt.Errorf("%d secret(s) out of sync", len(mismatches))
		}

		fmt.Printf("✓ All %d secret(s) in sync\n", processed)
	case constant.Encode:
		fmt.Printf(
			"\nEncoded %d secret(s), %d in sync\n",
			processed-inSync,
			inSync,
		)
	default:
		fmt.Printf("\nProcessed %d secret(s)\n", processed)
	}

	if len(fails) > 0 {
		return fmt.Errorf("%d secret(s) failed", len(fails))
	}

	return nil
}
