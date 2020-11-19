package stack

// CspStack is a struct with methods needed to implement the Stack interface.
type CspStack struct {
}

// NewCspStack returns an empty CspStack.
func NewCspStack() *CspStack {
	return &CspStack{}
}

// Size returns the size of the stack.
func (cs *CspStack) Size() int {
	return -1
}

// Push pushes value onto the stack.
func (cs *CspStack) Push(value interface{}) {

}

// Pop pops the value at the top of the stack and returns it.
func (cs *CspStack) Pop() (value interface{}) {
	return nil
}

func (cs *CspStack) run() {

}
