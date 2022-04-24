package factorial

import (
	"math"
	"math/big"
)

type HugeNumber struct {
	precision uint
	val       *big.Float
}

func NewHugeNumber() *HugeNumber {
	return &HugeNumber{val: big.NewFloat(0), precision: 64}
}

func (n *HugeNumber) SetInt64(val int64) {
	if n != nil {
		n.val = new(big.Float).SetInt64(val)
	}
}

func (n *HugeNumber) Value() *big.Float {
	if n != nil {
		return n.val
	}
	return NewHugeNumber().Value()
}

func (n *HugeNumber) ValueInt() *big.Int {
	if n != nil {
		s := n.Value().String()
		v, ok := new(big.Int).SetString(s, 10)
		if ok {
			return v
		}
	}
	return NewHugeNumber().ValueInt()
}

func (n *HugeNumber) SetValue(val string) {
	if n != nil {
		v, ok := new(big.Float).SetPrec(n.Prec()).SetString(val)
		if ok {
			n.val = v
		}
	}
}

func (n *HugeNumber) String() string {
	if n != nil {
		return n.Value().Text('g', n.PrecInt())
	}
	return NewHugeNumber().Value().Text('g', n.PrecInt())
}

func (n *HugeNumber) Prec() uint {
	if n != nil {
		return n.precision
	}
	return NewHugeNumber().precision
}

func (n *HugeNumber) PrecInt() int {
	return int(n.Prec())
}

func (n *HugeNumber) SetPrec(p uint) {
	if n != nil {
		n.precision = p
	}
}

// http://hyperphysics.phy-astr.gsu.edu/hbase/Math/stirling.html
// ln N! = NlnN - N
func (n *HugeNumber) factorialApproximationNlnN() *HugeNumber {
	if n != nil {
		N := n.Value()
		x, _ := n.Value().Float64()
		lnN := new(big.Float).SetPrec(n.Prec()).SetFloat64(math.Log(x))
		temp := new(big.Float)
		NlnN := temp.Mul(N, lnN)
		NlnN_N := temp.Sub(NlnN, N)
		n.SetValue(NlnN_N.Text('g', n.PrecInt()))
		return n
	}
	return NewHugeNumber()
}

func (n *HugeNumber) Factorial(p uint) string {
	if n != nil {
		n.SetPrec(p)
		n.factorialApproximationNlnN()
		exp := "e^" // math.E // 2.718281828459045
		return exp + n.String()
	}
	return ""
}
