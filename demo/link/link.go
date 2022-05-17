package link

import _ "unsafe"

//go:linkname helloWorld demo/hello.Hello
func helloWorld() {
	println("hello world!")
}
