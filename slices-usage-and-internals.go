package main

import "fmt"

func main() {

	//b := [2]string{"Penn", "Teller"}
	b := [...]string{"Penn", "Teller"}

	fmt.Printf("%v\n", b)
	fmt.Printf("len(b) = %d\n", len(b))

	letters := []string{"a", "b", "c", "d"}

	fmt.Printf("%v\n", letters)
	fmt.Printf("len(letters) = %d\n", len(letters))
}
