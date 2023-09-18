package db

import "ghoast/m/v2/pkg/xd"

func hi_test_hello() {
	println("Hello, World!")
	chained := xd.New()
	chained.ThisIsExported().ThisIsExported().ThisIsExported().ThisIsExported()
}
