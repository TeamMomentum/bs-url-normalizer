from ctypes import *
lib = cdll.LoadLibrary("./libmomentum_url_normalizer.a")
ptr = c_char_p("")
lib.first_normalize_url("http://example.com/path/", pointer(ptr))
print ptr
lib.free_normalize_url(pointer(ptr))

#ptr = lib.second_normalize_url(c_char_p("http://example.com/path/"))
#print c_char_p(ptr).value
#lib.free_normalize_url(ptr)
