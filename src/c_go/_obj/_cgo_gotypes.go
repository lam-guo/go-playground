// Code generated by cmd/cgo; DO NOT EDIT.

package main

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype__GoString_ string

type _Ctype_char int8

type _Ctype_intgo = _Ctype_ptrdiff_t

type _Ctype_long int64

type _Ctype_ptrdiff_t = _Ctype_long

//go:notinheap
type _Ctype_void_notinheap struct{}

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})


//go:linkname _cgo_runtime_gostring runtime.gostring
func _cgo_runtime_gostring(*_Ctype_char) string

// GoString converts the C string p into a Go string.
func _Cfunc_GoString(p *_Ctype_char) string {
	return _cgo_runtime_gostring(p)
}
//go:cgo_export_dynamic SayHello
//go:linkname _cgoexp_49c04f75d06d_SayHello _cgoexp_49c04f75d06d_SayHello
//go:cgo_export_static _cgoexp_49c04f75d06d_SayHello
func _cgoexp_49c04f75d06d_SayHello(a *struct {
		p0 *_Ctype_char
	}) {
	SayHello(a.p0)
}
