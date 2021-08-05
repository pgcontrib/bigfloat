package bigfloat_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pgcontrib/bigfloat"
)

func TestNewFloat(t *testing.T) {

	f1 := big.NewFloat(1000.01)
	f2 := big.NewFloat(1000.02)
	diff := big.NewFloat(0.01)
	// this test fails,
	t.Run("test equality", func(t *testing.T) {
		a := bigfloat.NewFloat(f1)
		b := bigfloat.NewFloat(f2)
		c := bigfloat.NewFloat(diff)

		d := b.Sub(a)

		assert.Equal(t, d.String(), c.String())

	})
}
