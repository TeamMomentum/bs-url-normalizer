#please make shared library first
from ctypes import *
lib = cdll.LoadLibrary("../../libmomentum_url_normalizer.a")

ptr = c_char_p("".encode())
lib.first_normalize_url("http://example.com/path/".encode(), pointer(ptr))
print(ptr.value.decode())
lib.free_normalize_url(pointer(ptr))

ptr = c_char_p("".encode())
lib.second_normalize_url("http://example.com/path/".encode(), pointer(ptr))
print(ptr.value.decode())
lib.free_normalize_url(pointer(ptr))
