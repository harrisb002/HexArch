// Will have its own adapter type (Same as all adapters)
package segment

type Adapter struct {
}

// Create an Adapter
func NewAdapter() *Adapter {
	// Inst. of Adapter Struct
	return &Adapter{}
}

// Implements the Segment Port interface
func (seg Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}
func (seg Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}
func (seg Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
func (seg Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}
