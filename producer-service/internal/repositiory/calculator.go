package repositiory

import "errors"

type Variables struct {
	a float64
	b float64
}

type Calculator interface {
	Multiply() float64
	Addition() float64
}

func NewVariables(a float64, b float64) (Calculator, error) {
	if a < 1 || b < 1 {
		return &Variables{}, errors.New("variables should be larger than zero")
	}
	return &Variables{
		a: a,
		b: b,
	}, nil
}

func (v *Variables) Multiply() float64 {
	return v.a * v.b
}

func (v *Variables) Addition() float64 {
	return v.a + v.b
}
