package factorial

type Calculater interface {
	SetInt64(v int64)
	Factorial(prec uint) string
}

var prec uint = 64

type Counter struct{}

func NewCounter() *Counter {
	return &Counter{}
}

func makeCalculater(val int64) Calculater {
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

func (s *Counter) Calculate(value int64) string {
	c := makeCalculater(value)
	if c != nil {
		return c.Factorial(Prec())
	}
	return "NaN"
}

func SetPrec(p uint) {
	prec = p
}

func Prec() uint {
	return prec
}
