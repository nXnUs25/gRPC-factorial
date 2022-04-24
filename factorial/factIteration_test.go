package factorial

import (
	"math"
	"math/big"
	"testing"
)

func TestFactorialIteraion(t *testing.T) {

	var prec uint = 30
	var bitSize int = 64

	tests := []struct {
		num  int64
		want *Number
	}{
		{
			num: 30,
			want: &Number{
				val:       265252859812191058636308480000000,
				precision: prec,
				bitSize:   bitSize,
			},
		},
		{
			num: 150,
			want: &Number{
				val:       5.7133839564458 * math.Pow10(262),
				precision: prec,
				bitSize:   bitSize,
			},
		},
		{
			num: 170,
			want: &Number{
				val:       7.257415615307899 * math.Pow10(306),
				precision: prec,
				bitSize:   bitSize,
			},
		},
		{
			num: 80,
			want: &Number{
				val:       7.15694570462638505039051537884 * math.Pow10(118),
				precision: prec,
				bitSize:   bitSize,
			},
		},
		{
			num: 5,
			want: &Number{
				val:       120,
				precision: prec,
				bitSize:   bitSize,
			},
		},
	}

	for _, tt := range tests {
		num1 := NewNumber()
		num1.SetPrec(prec)
		num1.SetInt64(tt.num)
		got := num1.factorialIteraion()
		tolerance := new(big.Float).SetPrec(prec).SetFloat64(float64(math.Pow10(-int(prec))))

		bigWant := new(big.Float).SetPrec(prec).SetFloat64(tt.want.Value())
		bigGot := new(big.Float).SetPrec(prec).SetFloat64(got.Value())

		diff := bigWant.Sub(bigWant, bigGot)
		diffAbs := diff.Abs(diff)

		if err := diffAbs.Cmp(tolerance); err > 0 {
			t.Errorf("factorial iteration of %+v want %+v, got %+v", tt.num, tt.want, got)
		}
	}

}
