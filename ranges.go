package main

import "fmt"

func main() {

	//using range to sum the elements of an array
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//iterate returning index and value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//iterate over key value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//iterage over just the keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//iterate over Unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
