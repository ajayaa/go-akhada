package main

import "fmt"

func foo(str string) {
	fmt.Println(&str)
}
func main() {
	//fmt.Println(syntaxError("ParseBool", "aj"))
	str := "ajaya"
	fmt.Println(&str)
	foo(str[0:2])
}
