package main

import (
	"io"
	"os"
	"fmt"
)

func main() {
	CopyFile("foo.txt", "bar.txt")
	//a()
	//b()
	fmt.Println(c())
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++ 
	return
}

func b() {
	for i:=0; i<4; i++ {
		defer fmt.Println(i)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}


func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
