package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := &Product{
		Price: 1.00,
	}
	//err := p.Validate()
	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}
