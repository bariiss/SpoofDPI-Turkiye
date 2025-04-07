package version

import _ "embed"

//go:embed VERSION
var VERSION string

func PrintVersion() {
	println("spoofdpi", "v"+VERSION)
	println("Spoof DPI'ın bu sürümü Türkiye'de kullanılmak üzere yapılandırılmıştır.")
	println("https://github.com/bariiss/SpoofDPI-Turkiye")
}
