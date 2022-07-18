package repositiory

import "errors"

type Variables struct {
	a int
	b int
}

type Calculator interface {
	Multiply() int
	Addition() int
}

func NewVariables(a int, b int) (Calculator, error) {
	if a < 1 || b < 1 {
		return &Variables{}, errors.New("variables should be larger than zero")
	}
	return &Variables{
		a: a,
		b: b,
	}, nil
}

func (v *Variables) Multiply() int {
	return v.a * v.b
}

func (v *Variables) Addition() int {
	return v.a + v.b
}
