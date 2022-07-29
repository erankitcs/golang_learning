package data

import "testing"

func TestProductValidate(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 1.00,
		SKU:   "sss-vggvgv-vttv",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
