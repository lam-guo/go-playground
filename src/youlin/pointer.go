package youlin

import "unsafe"

// go:noinline
func convert(p *int) {
	q := (*int32)(unsafe.Pointer(p))
	*q = 0
}
