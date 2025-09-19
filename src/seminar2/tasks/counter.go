package tasks

type counter struct {
}

func newCounter(_ int) *counter {
	return &counter{}
}

func (c *counter) Increment() {
}

func (c *counter) Decrement() {
}

func (c *counter) GetValue() int {
	return 0
}

func (c *counter) Reset() {
}

func (c *counter) Add(_ int) {
}

func (c *counter) Subtract(_ int) {
}
