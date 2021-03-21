package main

import "fmt"

//输出小写得a-z和大写得Z-A
func main() {
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", 'a'+i)

	}
	fmt.Println()
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", 'Z'-i)

	}
}
