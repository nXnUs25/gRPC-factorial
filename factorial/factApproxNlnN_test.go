package factorial

import (
	"math"
	"math/big"
	"testing"
)

func TestFactorialApproximationNlnN(t *testing.T) {

	var prec uint = 30

	tests := []struct {
		num  int64
		want string
	}{
		{
			num:  5000000000,
			want: "1.066635187e+11",
		},

		{
			num:  10000,
			want: "82103.40369",
		},
		{
			num:  5000000000,
			want: "1.066635187e+11",
		},
		{
			num:  9999999999999,
			want: "2.893360619e+14",
		},
	}

	for _, tt := range tests {
		num1 := NewHugeNumber()
		num1.SetPrec(prec)
		num1.SetInt64(tt.num)
		got := num1.factorialApproximationNlnN()
		tolerance := new(big.Float).SetPrec(prec).SetFloat64(float64(math.Pow10(-int(prec))))

		bigWant, ok := new(big.Float).SetPrec(prec).SetString(tt.want)
		if !ok {
			t.Errorf("cannot parse %v", tt.want)
		}

		diff := bigWant.Sub(bigWant, got.Value())
		diffAbs := diff.Abs(diff)

		if err := diffAbs.Cmp(tolerance); err > 0 {
			t.Errorf("factorial approximation NlnN (Sterlinga) of %+v want %+v, got %+v", tt.num, tt.want, got)
		}
		t.Logf("[PASS] Passed with tolerance [%+v] for values [got: %v - want: %v]", tolerance, got, tt.want)
	}

}
