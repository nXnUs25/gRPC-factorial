package factorial

import (
	"math/big"
)

type BigNumber struct {
	precision uint
	val       *big.Float
}

func NewBigNumber() *BigNumber {
	return &BigNumber{val: big.NewFloat(0), precision: 64}
}

func (n *BigNumber) SetInt64(val int64) {
	if n != nil {
		n.val = new(big.Float).SetInt64(val)
	}
}

func (n *BigNumber) Value() *big.Float {
	if n != nil {
		return n.val
	}
	return NewBigNumber().Value()
}

func (n *BigNumber) ValueInt() *big.Int {
	if n != nil {
		s := n.Value().String()
		v, ok := new(big.Int).SetString(s, 10)
		if ok {
			return v
		}
	}
	return NewBigNumber().ValueInt()
}

func (n *BigNumber) SetValue(val string) {
	if n != nil {
		v, ok := new(big.Float).SetPrec(n.Prec()).SetString(val)
		if ok {
			n.val = v
		}
	}
}

func (n *BigNumber) String() string {
	if n != nil {
		return n.Value().Text('g', n.PrecInt())
	}
	return NewBigNumber().Value().Text('g', n.PrecInt())
}

func (n *BigNumber) Prec() uint {
	if n != nil {
		return n.precision
	}
	return NewBigNumber().precision
}

func (n *BigNumber) PrecInt() int {
	return int(n.Prec())
}

func (n *BigNumber) SetPrec(p uint) {
	if n != nil {
		n.precision = p
	}
}

func (n *BigNumber) factorialBigMul() *BigNumber {
	if n != nil {
		var cal = new(big.Int)
		f := n.ValueInt().Int64()
		bigNum := cal.MulRange(1, f)
		n.SetValue(bigNum.String())
		return n
	}
	return NewBigNumber()
}

func (n *BigNumber) Factorial(p uint) string {
	if n != nil {
		n.SetPrec(p)
		n.factorialBigMul()
		return n.String()
	}
	return ""
}
