package main

import "fmt"

// error ,just test for local value
func main() {
	var a = 12
	if a > 10 {
		b := 3
		fmt.Println(b)
		var c = 5
	}
	fmt.Println(c)

	fmt.Println(b)

}
