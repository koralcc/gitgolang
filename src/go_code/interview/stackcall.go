package main

func main() {
	caller()
}

func caller() {
	var a int64 = 1
	var b int64 = 2
	callee(a, b)
}

func callee(a, b int64) (int64, int64) {
	c := a + 5
	d := b * 4
	return c, d
}
