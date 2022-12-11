package musgo_int

import (
	"testing"
)

func TestMusgoInt(t *testing.T) {
	var n int = 5
	buf := make([]byte, Int(n).SizeMUS())
	Int(n).MarshalMUS(buf)

	var an int
	(*Int)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
