package youlin

//go:noinline
func New() int {
	p := new(int)
	return *p
}

//go:noinline
func NewInt() *int {
	var a int
	return &a
}
