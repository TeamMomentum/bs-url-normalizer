import com.sun.jna
import com.sun.jna.ptr.PointerByReference

trait NormURL extends jna.Library {
  def first_normalize_url(rawurl: String, ptr: PointerByReference)
  def second_normalize_url(rawurl: String, ptr: PointerByReference)
  def free_normalize_url(ptr: PointerByReference)
}

object Sample {
  val normURL = jna.Native.loadLibrary("momentum_url_normalizer", classOf[NormURL]).asInstanceOf[NormURL]

  private def getNormalizedURL(url: String, normalizeURL: (String, PointerByReference) => Unit): String = {
    val ptr = new PointerByReference
    normalizeURL(url, ptr)  // URL正規化処理
    val normalized = ptr.getValue.getString(0)  // 正規化済みURL文字列の取得
    normURL.free_normalize_url(ptr) // 使用済みメモリの解放
    normalized
  }

  def getFirstNormalizedURL(url: String): String = getNormalizedURL(url, normURL.first_normalize_url)
  def getSecondNormalizedURL(url: String): String = getNormalizedURL(url, normURL.second_normalize_url)

  def main(args: Array[String]) = {
    // 第一段階正規化
    val fnu = getFirstNormalizedURL("http://example.com/tihoukoumu?d=1&a=2&c=3&b=4")
    println("First normalized URL: " + fnu)

    // 第二段階正規化
    val snu = getSecondNormalizedURL("http://example.com/tihoukoumu?d=1&a=2&c=3&b=4")
    println("Second normalized URL: " + snu)
  }
}
