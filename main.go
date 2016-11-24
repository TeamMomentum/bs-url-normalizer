package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"net/url"
	"unsafe"

	"github.com/TeamMomentum/bscore/lib/utils/urls"
)

//export first_normalize_url
func first_normalize_url(raw *C.char, ptr *unsafe.Pointer) {
	goraw := C.GoString(raw)
	ul, err := url.Parse(goraw)
	if err != nil {
		*ptr = unsafe.Pointer(C.CString(""))
	}
	*ptr = unsafe.Pointer(C.CString(urls.FirstNormalizeURL(ul)))
}

//export second_normalize_url
func second_normalize_url(raw *C.char, ptr *unsafe.Pointer) {
	goraw := C.GoString(raw)
	ul, err := url.Parse(goraw)
	if err != nil {
		*ptr = unsafe.Pointer(C.CString(""))
	}
	*ptr = unsafe.Pointer(C.CString(urls.SecondNormalizeURL(ul)))

}

//export free_normalize_url
func free_normalize_url(mem unsafe.Pointer) {
	C.free(mem)
}

func main() {
}
