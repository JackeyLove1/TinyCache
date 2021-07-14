package TinyCache

type ByteView struct {
	b []byte
}

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// Returns a copy of the data of a byte slice
func (v ByteView) ByteSlice() []byte {
	c := make([]byte, v.Len())
	copy(c, v.b)
	return c
}

// String returns the data as a string
func (v ByteView) String() string {
	return string(v.b)
}
