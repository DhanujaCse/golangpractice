package arrayws

type MyInterface interface {
	Add() int
	Sub() int
}

type CalculateAdd struct {
	Num1 int
	Num2 int
}

func (a CalculateAdd) Add() int {
	return a.Num1 + a.Num2

}
func (s CalculateAdd) Sub() int {
	return s.Num1 - s.Num2
}

type Multiply interface {
	Multiply1() int
	Divide() int
}
type DigitsNum struct {
	Number1 int
	Number2 int
}

func (d DigitsNum) Multiply1() int {
	return d.Number1 * d.Number2

}
func (d DigitsNum) Divide() int {
	return d.Number1 / d.Number2
}
