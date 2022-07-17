package repositiory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatorMultiply(t *testing.T) {
	type testCase struct {
		name   string
		a      float64
		b      float64
		result float64
	}

	testCases := []testCase{
		{
			name:   "Valid values",
			a:      1.00,
			b:      2.00,
			result: 2,
		},
		{
			name:   "Invalid values",
			a:      4.00,
			b:      2.00,
			result: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calculator, err := NewVariables(tc.a, tc.b)

			result := calculator.Multiply()

			// Assert
			assert.Equal(t, tc.result, result)
			assert.NoError(t, err)
		})
	}
}

func TestCalculatorMultiplyError(t *testing.T) {
	variables, err := NewVariables(0.0, 1.0)

	assert.Equal(t, &Variables{}, variables)
	assert.Error(t, err)
}

func TestCalculatorAddition(t *testing.T) {
	variables, err := NewVariables(1.0, 1.0)

	result := variables.Addition()

	assert.Equal(t, 2.0, result)
	assert.NoError(t, err)
}
