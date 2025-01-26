package employee

type Employee struct {
	AgeInt int
}

func New(age int) *Employee {
	e := Employee{AgeInt: age}
	return &e
}

func (e *Employee) Age() int {
	return e.AgeInt
}
