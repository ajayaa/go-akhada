package main

import "fmt"
import "strconv"

//import "syscall"
import "os"

func main() {
	str := "qw123"
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		fmt.Println(err)
		//syscall.Exit(1)
		os.Exit(1)
	}
	fmt.Println(i)
}
