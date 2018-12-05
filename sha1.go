package main

import "crypto/sha1"
import "crypto/md5"
import "fmt"

func main() {

	s := "sha1 this string"

	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	h2 := md5.New()
	h2.Write([]byte(s))
	bs2 := h2.Sum(nil)

	fmt.Println(s)

	fmt.Printf("Sha1: %x\n", bs)
	fmt.Printf("MD5: %x\n", bs2)
}
