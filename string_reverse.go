package main

import "fmt"

func reverse_word_by_word(str string) string {
	length := len(str)
	s := []rune(str)
	i := 0
	j := 0
	for ; j < length; j++ {
		if s[j] == ' ' {
			reverse_rune(s, i, j-1)
			i = j + 1
		}
	}
	reverse_rune(s, i, j-1)
	return string(s)
}

func reverse(str string) string {
	len := len(str)
	s := []rune(str)
	reverse_rune(s, 0, len-1)
	return string(s)
}

func reverse_rune(s []rune, i int, j int) {
	mid := (j - i) / 2
	for k := 0; k < mid; k++ {
		temp := s[i+k]
		s[i+k] = s[j-k]
		s[j-k] = temp
	}
}

func change_arr(arr []int) {
	arr = []int{-1, -2, -3}
	arr[0] = -1
}

func main() {
	//fmt.Println("Hi there!")
	str := "  abc  def  "
	fmt.Println(reverse_word_by_word(str))
	//s := []rune(str)
	//fmt.Println(s[3] == ' ')
	//arr := []int{1, 2, 3}
	//fmt.Println(arr)
	//change_arr(arr)
	//fmt.Println(arr)
}
