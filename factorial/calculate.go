package factorial

type Calculater interface {
	SetInt64(v int64)
	Factorial(prec uint) string
}

func MakeCalculate(val int64) Calculater {
	var c Calculater
	switch {
	// 170 its the last value which can be calculated 171 gives +inf
	case val <= 170:
		c = NewNumber()
		c.SetInt64(val)
		return c
	// seems with 100000 i am getting most time wise reasonable results
	case 170 < val && val <= 100000:
		c = NewBigNumber()
		c.SetInt64(val)
		return c
	default:
		c = NewHugeNumber()
		c.SetInt64(val)
		return c
	}
}

func Calculate(c Calculater, prec uint) string {
	if c != nil {
		return c.Factorial(prec)
	}
	return "NaN"
}
