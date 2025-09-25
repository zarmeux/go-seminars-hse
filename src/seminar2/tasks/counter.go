package tasks

type counter struct {
	value int
}

func newCounter(new_value int) *counter {
	return &counter{value: new_value}
}

func (c *counter) Increment() {
	c.value++
}

func (c *counter) Decrement() {
	c.value--
}

func (c *counter) GetValue() int {
	return c.value
}

func (c *counter) Reset() {
	c.value = 0
}

func (c *counter) Add(delta int) {
	c.value += delta
}

func (c *counter) Subtract(delta int) {
	c.value -= delta
}
