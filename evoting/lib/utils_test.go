package lib

import (
	"testing"

	"github.com/dedis/kyber"

	"github.com/stretchr/testify/assert"
)

func TestDKGSimulate(t *testing.T) {
	dkgs, _ := DKGSimulate(5, 4)
	assert.Equal(t, 5, len(dkgs))

	secrets := make([]*SharedSecret, 5)
	for i, dkg := range dkgs {
		secrets[i], _ = NewSharedSecret(dkg)
	}

	var private kyber.Scalar
	for _, secret := range secrets {
		if private != nil {
			assert.NotEqual(t, private.String(), secret.V.String())
		}
		private = secret.V
	}
}
