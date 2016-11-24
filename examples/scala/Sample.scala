import com.sun.jna
import com.sun.jna.ptr.PointerByReference

trait NormURL extends jna.Library {
  def first_normalize_url(rawurl: String, ptr: PointerByReference)
  def second_normalize_url(rawurl: String,  ptr: PointerByReference)
  def free_normalize_url(ptr: PointerByReference)
}

val normURL = jna.Native.loadLibrary("normurl", classOf[NormURL]).asInstanceOf[NormURL]

// 第一段階正規化
val ptr1 = new PointerByReference
normURL.first_normalize_url("http://example.com/tihoukoumu?d=1&a=2&c=3&b=4", ptr1)
val fnu: String = ptr1.getValue.getString(0) // 正規化済みURL文字列の取得
normURL.free_normalize_url(ptr1) // 使用済みメモリの解放

// 第二段階正規化
val ptr2 = new PointerByReference
normURL.second_normalize_url("http://example.com/tihoukoumu?d=1&a=2&c=3&b=4", ptr2)
val snu: String = ptr2.getValue.getString(0) // 正規化済みURL文字列の取得
normURL.free_normalize_url(ptr2) // 使用済みメモリの解放
