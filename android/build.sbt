name := "duckspeak-android"

enablePlugins(AndroidApp)
android.useSupportVectors

versionCode := Some(1)
version := "0.1-SNAPSHOT"

instrumentTestRunner :=
  "android.support.test.runner.AndroidJUnitRunner"

platformTarget := "android-24"
minSdkVersion := "19"

scalaVersion := "2.11.8"
scalacOptions ++= Seq(
  "-deprecation",
  "-feature",
  "-Xfatal-warnings"
)

updateCheck in Android := {}

proguardCache in Android ++= Seq("org.scaloid")
proguardOptions in Android ++= Seq(
  "-dontobfuscate",
  "-dontoptimize",
  "-keepattributes Signature",
  "-printseeds target/seeds.txt",
  "-printusage target/usage.txt",
  "-dontwarn scala.collection.**",
  "-dontwarn org.scaloid.**"
)

javacOptions in Compile ++= Seq(
  "-source", "1.7",
  "-target", "1.7"
)

libraryDependencies ++= Seq(
  "org.scaloid" %% "scaloid" % "4.2",

  // java deps
  aar("com.android.support" % "appcompat-v7" % "24.0.0"),
  "com.android.support.test" % "runner" % "0.5" % "androidTest",
  "com.android.support.test.espresso" % "espresso-core" % "2.2.2" % "androidTest"
)

run <<= run in Android
install <<= install in Android

unmanagedClasspath in Test ++= (bootClasspath in Android).value