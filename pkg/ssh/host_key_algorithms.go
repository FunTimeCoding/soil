package ssh

import (
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	soilErrors "github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"net"
)

func hostKeyAlgorithms(
	callback ssh.HostKeyCallback,
	address string,
) []string {
	public, _, e := ed25519.GenerateKey(rand.Reader)
	soilErrors.PanicOnError(e)
	probe, f := ssh.NewPublicKey(public)
	soilErrors.PanicOnError(f)
	var mismatch *knownhosts.KeyError

	if !errors.As(
		callback(address, &net.TCPAddr{IP: net.IPv4zero}, probe),
		&mismatch,
	) {
		return nil
	}

	var result []string

	for _, k := range mismatch.Want {
		if k.Key.Type() == ssh.KeyAlgoRSA {
			result = append(result, ssh.KeyAlgoRSASHA512, ssh.KeyAlgoRSASHA256)
		}

		result = append(result, k.Key.Type())
	}

	return result
}
