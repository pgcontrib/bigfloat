package bigfloat_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pgcontrib/bigfloat"
)

func TestNewFloat(t *testing.T) {

	f1 := big.NewFloat(1000.01)
	f2 := big.NewFloat(1000.02)
	diff := big.NewFloat(0.01)

	t.Run("test equality", func(t *testing.T) {
		a := bigfloat.NewFloat(f1)
		b := bigfloat.NewFloat(f2)
		c := bigfloat.NewFloat(diff)

		d := b.Sub(a)
		fmt.Printf("%v - %v  = %v \n", b, a, d)
		assert.Equal(t, d, c)

	})
}
