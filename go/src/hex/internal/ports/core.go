package ports

// The core adapter must comply by this interface.
// Simplified as of now
type Segment interface {
	Addition(a int32, b int32) (int32, error)
	Subtraction(a int32, b int32) (int32, error)
	Division(a int32, b int32) (int32, error)
	Multiplication(a int32, b int32) (int32, error)
}
