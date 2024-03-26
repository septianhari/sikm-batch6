package internal

type Calculator struct {
	Base float32
}

func NewCalculator(base float32) *Calculator {
	return &Calculator{Base: base}
}

func (c *Calculator) Add(num float32) {
	// TODO: answer here
}

func (c *Calculator) Subtract(num float32) {
	// TODO: answer here
}

func (c *Calculator) Multiply(num float32) {
	// TODO: answer here

}

func (c *Calculator) Divide(num float32) {
	// TODO: answer here
}

func (c *Calculator) Result() float32 {
	return 0 // TODO: replace this
}
