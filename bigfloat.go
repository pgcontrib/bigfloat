package bigfloat

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"math/big"
)

type (
	Bigfloat big.Float
)

func NewFloat(f *big.Float) *Bigfloat {
	return (*Bigfloat)(f)
}

func (f *Bigfloat) Value() (driver.Value, error) {
	return (*big.Float)(f).String(), nil
}

func (f *Bigfloat) Scan(value interface{}) error {

	var i sql.NullString

	if err := i.Scan(value); err != nil {
		return err
	}

	if _, ok := (*big.Float)(f).SetString(i.String); ok {
		return nil
	}

	return fmt.Errorf("Error converting type %T into Bigint", value)
}

func (f *Bigfloat) toBigFloat() *big.Float {
	return (*big.Float)(f)
}

func (f *Bigfloat) Add(target *Bigfloat) *Bigfloat {
	return NewFloat(new(big.Float).Add(f.toBigFloat(), target.toBigFloat()))
}
func (f *Bigfloat) Sub(target *Bigfloat) *Bigfloat {
	return NewFloat(new(big.Float).Sub(f.toBigFloat(), target.toBigFloat()))
}
func (f *Bigfloat) Mul(target *Bigfloat) *Bigfloat {
	return NewFloat(new(big.Float).Mul(f.toBigFloat(), target.toBigFloat()))
}
func (f *Bigfloat) Div(target *Bigfloat) *Bigfloat {
	return NewFloat(new(big.Float).Quo(f.toBigFloat(), target.toBigFloat()))
}
func (f *Bigfloat) Neg() *Bigfloat {
	return NewFloat(new(big.Float).Neg(f.toBigFloat()))
}
func (f *Bigfloat) Abs() *Bigfloat {
	return NewFloat(new(big.Float).Abs(f.toBigFloat()))
}
func (f *Bigfloat) Cmp(target *Bigfloat) int {
	return (f.toBigFloat().Cmp(target.toBigFloat()))
}
func (f *Bigfloat) String() string {
	return (f.toBigFloat()).String()
}
func (f *Bigfloat) ToFloat64() (float64, int8) {
	x, a := f.toBigFloat().Float64()
	return x, int8(a)
}

func FromString(inp string) (*Bigfloat, error) {
	f, _, err := big.ParseFloat(inp, 10, 10, big.ToZero)

	if err != nil {
		return nil, err
	}
	return NewFloat(f), nil
}
