package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"net/url"
	"unsafe"

	"./lib/urls"
)

//export first_normalize_url
func first_normalize_url(cStr *C.char, p *unsafe.Pointer) {
	rawURL := C.GoString(cStr)
	ul, err := url.ParseRequestURI(rawURL)
	if err != nil {
		*p = unsafe.Pointer(C.CString(""))
	} else {
		*p = unsafe.Pointer(C.CString(urls.FirstNormalizeURL(ul)))
	}
}

//export second_normalize_url
func second_normalize_url(cStr *C.char, p *unsafe.Pointer) {
	rawURL := C.GoString(cStr)
	ul, err := url.ParseRequestURI(rawURL)
	if err != nil {
		*p = unsafe.Pointer(C.CString(""))
	} else {
		*p = unsafe.Pointer(C.CString(urls.SecondNormalizeURL(ul)))
	}
}

//export free_normalize_url
func free_normalize_url(m *unsafe.Pointer) {
	if m != nil { // check for avoiding panic
		C.free(*m)
	}
}

func main() {
}
