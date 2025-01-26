package customer

type Customer struct {
	AgeInt int
}

func New(age int) *Customer {
	c := Customer{AgeInt: age}
	return &c
}

func (c *Customer) Age() int {
	return c.AgeInt
}
