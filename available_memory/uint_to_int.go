package main

import "fmt"

//import "math"

func main() {
	var a uint64
	a = 20
	var b int64
	b = int64(a)
	fmt.Printf("%T\n", b)
	//fmt.Printf("%T\n", math.Min(a, b))
	defaultCgroupMemPath := "/usr/local/lib"
	fmt.Printf("can't undestand format of %s, setting rocksdb cache size to 512MB", defaultCgroupMemPath)
}
