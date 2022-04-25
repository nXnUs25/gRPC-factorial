package factorial

import (
	"math/big"
	"reflect"
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
		got := MakeCalculate(tt.num)

		if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
			t.Errorf("InitCalculate type error got %v, want %v", got, tt.want)
		}
		t.Logf("[PASS] Initialized Types are as expected got: [%v] and want it [%v]", reflect.TypeOf(got), reflect.TypeOf(tt.want))
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
	hugeNum, _ := new(big.Float).SetPrec(prec).SetString("1.066635187e+11")
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
		c := MakeCalculate(tt.num)
		got := Calculate(c, prec)
		if got != tt.want {
			t.Errorf("Wrong value of Calculate for %+v want %+v, got %+v", tt.num, tt.want, got)
		}
		t.Logf("[PASS] Values calculated [%s] and want it [%s] are same.", tt.want, got)
	}
}
