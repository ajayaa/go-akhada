package main

import "fmt"
import "bytes"

func print_str(str string) {
	fmt.Println(str)
}
func main() {
	var buf bytes.Buffer
	buf.WriteString("a")
	buf.WriteString("b")
	fmt.Println(buf.String())
	var str string
	str += "aj"
	fmt.Println(str == "")
	print_str(str)
}
