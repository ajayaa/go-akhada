package main

import (
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS == "linux" {
		fmt.Println(runtime.GOOS)
		fmt.Printf("%T\n", runtime.GOOS)
	}
}
