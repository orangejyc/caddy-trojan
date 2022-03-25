//go:build malloc_cgo
// +build malloc_cgo

package memory

// #include <stdlib.h>
import "C"

import "unsafe"

// Alloc is ...
func Alloc[T any](n int) Array[T] {
	return Array[T]{data: unsafe.Slice((*T)(unsafe.Pointer(uintptr(C.malloc(C.size_t(n*int(unsafe.Sizeof(*(new(T))))))))), n)}
}

// Free is ...
func Free[T any](arr Array[T]) {
	C.free(unsafe.Pointer(&arr.data[0]))
}
