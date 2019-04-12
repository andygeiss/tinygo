package views

import "syscall/js"

//go:export HomeSayHello
func HomeSayHello() {
	js.Global().Call("alert", "alert called from Go!")
}
