scalaVersion := "2.11.8"

// https://mvnrepository.com/artifact/net.java.dev.jna/jna
libraryDependencies += "net.java.dev.jna" % "jna" % "4.2.2"

// [Optional] SharedLibrary ディレクトリの指定
unmanagedClasspath in Runtime += baseDirectory.value / "lib.shared"
