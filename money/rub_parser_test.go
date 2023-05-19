package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RubParse(t *testing.T) {
	parser := NewParser()
	res, err := parser.ParseFloat("50,00")
	assert.NoError(t, err)
	assert.Equal(t, 50.00, res)

	res, err = parser.ParseFloat("50,99")
	assert.NoError(t, err)
	assert.Equal(t, 50.99, res)

	res, err = parser.ParseFloat("0,10")
	assert.NoError(t, err)
	assert.Equal(t, 0.1, res)

	res, err = parser.ParseFloat("1.22")
	assert.NoError(t, err)
	assert.Equal(t, 1.22, res)
}
