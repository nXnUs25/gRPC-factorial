package factorial

import (
	"math"
	"math/big"
	"testing"
)

func TestFactorialBigMul(t *testing.T) {

	var prec uint = 30

	tests := []struct {
		num  int64
		want string
	}{
		{
			num:  5,
			want: "120",
		},
		{
			num:  1000,
			want: "4.02387260077093773543702433923e+2567",
		},
		{
			num:  10000,
			want: "2.846259680917054519e+35659",
		},
		{
			num:  1000000,
			want: "8.26393168833124006237664610317e+5565708",
		},
	}

	for _, tt := range tests {
		num1 := NewBigNumber()
		num1.SetPrec(prec)
		num1.SetInt64(tt.num)
		got := num1.factorialBigMul()
		tolerance := new(big.Float).SetPrec(prec).SetFloat64(float64(math.Pow10(-int(prec))))

		bigWant, ok := new(big.Float).SetPrec(prec).SetString(tt.want)
		if !ok {
			t.Errorf("[FAIL] Cannot parse %v", tt.want)
		}

		diff := bigWant.Sub(bigWant, got.Value())
		diffAbs := diff.Abs(diff)

		if err := diffAbs.Cmp(tolerance); err > 0 {
			t.Errorf("[FAIL] Factorial by math/big.RangeMul of %+v want %+v, got %+v", tt.num, tt.want, got)
		} else {
			t.Logf("[PASS] Passed with tolerance [%+v] for values [got: %v - want: %v]", tolerance, got, tt.want)
		}
	}

}
