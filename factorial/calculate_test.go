package factorial

import (
	"math"
	"math/big"
	"reflect"
	"strings"
	"testing"
)

func TestMakeCalculate(t *testing.T) {

	tests := []struct {
		num  int64
		want Calculater
	}{
		{
			num:  50,
			want: &Number{},
		},
		{
			num:  10000,
			want: &BigNumber{},
		},
		{
			num:  5000000000,
			want: &HugeNumber{},
		},
	}

	for _, tt := range tests {
		got := makeCalculater(tt.num)

		if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
			t.Errorf("[FAIL] InitCalculate type error got %v, want %v", got, tt.want)
		} else {
			t.Logf("[PASS] Initialized Types are as expected got: [%v] and want it [%v]", reflect.TypeOf(got), reflect.TypeOf(tt.want))
		}
	}
}

func TestCalculate(t *testing.T) {

	var (
		prec    uint = 30
		bitSize      = 64
	)

	bigNum, ok := new(big.Float).SetPrec(prec).SetString("2.846259680917054519e+35659")
	if !ok {
		t.Errorf("cannot parse %v", bigNum)
	}
	hugeNum, ok := new(big.Float).SetPrec(prec).SetString("1.066635187e+11")
	if !ok {
		t.Errorf("cannot parse %v", hugeNum)
	}

	n := &Number{
		val:       120,
		precision: prec,
		bitSize:   bitSize,
	}

	bn := &BigNumber{
		val:       bigNum,
		precision: prec,
	}

	hn := &HugeNumber{
		val:       hugeNum,
		precision: prec,
	}

	tests := []struct {
		num  int64
		want string
	}{
		{
			num:  5,
			want: n.String(),
		},
		{
			num:  10000,
			want: bn.String(),
		},
		{
			num:  5000000000,
			want: "e^" + hn.String(),
		},
	}

	for _, tt := range tests {
		c := NewCounter()
		SetPrec(prec)
		got := c.Calculate(tt.num)
		tolerance := new(big.Float).SetPrec(prec).SetFloat64(float64(math.Pow10(-int(prec))))

		bigWant, ok := new(big.Float).SetPrec(prec).SetString(strings.Trim(tt.want, "e^"))
		if !ok {
			t.Errorf("[FAIL] Cannot parse %v", tt.want)
		}

		bigGot, ok := new(big.Float).SetPrec(prec).SetString(strings.Trim(got, "e^"))
		if !ok {
			t.Errorf("[FAIL] Cannot parse %v", tt.want)
		}

		diff := bigWant.Sub(bigWant, bigGot).SetPrec(prec)
		diffAbs := diff.Abs(diff)

		if err := diffAbs.Cmp(tolerance); err > 0 {
			t.Errorf("[FAIL] Get value of Calculate with tolerance [%v] for [%v] want [%v], got [%v]", tolerance, tt.num, tt.want, got)
		} else {
			t.Logf("[PASS] Values calculated [%s] and want it [%s] are same.", tt.want, got)
		}
	}
}
