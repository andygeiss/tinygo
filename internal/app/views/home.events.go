package views

import "syscall/js"

//go:export homeSayHello
func HomeSayHello() {
	js.Global().Call("alert", "alert called from Go!")
}
