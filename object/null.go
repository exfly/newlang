package object

// Null is the null type and used to represent the absence of a value
type Null struct{}

func (n *Null) Bool() bool {
	return false
}

func (n *Null) Compare(other Object) int {
	if _, ok := other.(*Null); ok {
		return 0
	}
	return 1
}

func (n *Null) String() string {
	return n.Inspect()
}

// Type returns the type of the object
func (n *Null) Type() Type { return NULL }

// Inspect returns a stringified version of the object for debugging
func (n *Null) Inspect() string { return "null" }
