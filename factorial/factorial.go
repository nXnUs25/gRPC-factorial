package factorial

import (
	"strconv"
)

type Number struct {
	precision uint
	bitSize   int
	val       float64
}

func NewNumber() *Number {
	return &Number{val: 0, precision: 64, bitSize: 64}
}

func (n *Number) String() string {
	if n != nil {
		return strconv.FormatFloat(n.Value(), 'g', n.PrecInt(), n.bitSize)
	}
	return NewNumber().String()

}

func (n *Number) Value() float64 {
	if n != nil {
		return n.val
	}
	return NewNumber().Value()
}

func (n *Number) SetPrec(p uint) {
	if n != nil {
		n.precision = p
	}
}

func (n *Number) Prec() uint {
	if n != nil {
		return n.precision
	}
	return NewNumber().Prec()
}

func (n *Number) PrecInt() int {
	return int(n.Prec())
}

func (n *Number) SetValue(value float64) {
	if n != nil {
		n.val = value
	}
}

func (n *Number) SetInt64(value int64) {
	if n != nil {
		n.val = float64(value)
	}
}

func (n *Number) SetNumber32() {
	if n != nil {
		n.bitSize = 32
	}
}

func (n *Number) SetNumber64() {
	if n != nil {
		n.bitSize = 64
	}
}

// for small numbers, and using iteration as is better option
// faster, do not use the staxk memory and is more efficient than recursion
// it will be used for small number
func (n *Number) factorialIteraion() *Number {
	if n != nil {
		of := n.Value()
		if of > 0 {
			var i float64
			for i = 1; i < of; i++ {
				n.SetValue(i * n.Value())
			}
			return n
		}
	}
	n = NewNumber()
	n.SetValue(1)
	return n
}

func (n *Number) Factorial(p uint) string {
	if n != nil {
		n.SetPrec(p)
		n.factorialIteraion()
		return n.String()
	}
	return ""
}
