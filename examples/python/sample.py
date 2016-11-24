#please make shared library first
from ctypes import *
lib = cdll.LoadLibrary("../../libmomentum_url_normalizer.a")

ptr = c_char_p("")
lib.first_normalize_url("http://example.com/path/", pointer(ptr))
print ptr.value
lib.free_normalize_url(ptr)

ptr = c_char_p("")
lib.second_normalize_url("http://example.com/path/", pointer(ptr))
print ptr.value
lib.free_normalize_url(ptr)
