package examples.java;

import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.ptr.PointerByReference;

interface NormURL extends Library {
    NormURL INATANCE = (NormURL)Native.loadLibrary("momentum_url_normalizer", NormURL.class);   // expects libnormurl.so (Linux) or libnormurl.dylib (darwin), ...
    void first_normalize_url(String rawurl, PointerByReference ptr);
    void second_normalize_url(String rawurl,  PointerByReference ptr);
    void free_normalize_url(PointerByReference ptr);
}

class Sample {
    public static void main(String[] args) {
        final NormURL normURL = NormURL.INATANCE;

        // 第一段階正規化
        final PointerByReference ptr1 = new PointerByReference();
        normURL.first_normalize_url("http://example.com/tihoukoumu?d=1&a=2&c=3&b=4", ptr1);
        final String fnu = ptr1.getValue().getString(0); // 正規化済みURL文字列の取得
        System.out.println("First normalized URL: " + fnu);
        normURL.free_normalize_url(ptr1); // 使用済みメモリの解放
        // 第二段階正規化
        final PointerByReference ptr2 = new PointerByReference();
        normURL.second_normalize_url("http://example.com/tihoukoumu?d=1&a=2&c=3&b=4", ptr2);
        final String snu = ptr2.getValue().getString(0); // 正規化済みURL文字列の取得
        System.out.println("Second normalized URL: " + snu);
        normURL.free_normalize_url(ptr2); // 使用済みメモリの解放
        return;
    }
}
