package main

import "fmt"
import "os"

func main() {
	str, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting the hostname")
	}
	fmt.Fprintf(os.Stdout, str+"\n")
}
