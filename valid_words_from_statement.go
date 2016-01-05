package main

import "fmt"
import "io/ioutil"
import "strings"

//Reads the dictionary file in Linux system.
func ReadDict(path string) map[string]bool {
	m := make(map[string]bool)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	str := string(data)
	words := strings.Split(str, "\n")
	for _, word := range words {
		if len(word) != 1 {
			m[word] = true
		}
	}
	return m
}

func is_valid_string_valid_sentence(str string, m map[string]bool) bool {
	//fmt.Println(str)
	if _, ok := m[str]; ok {
		return true
	}
	len := len(str)
	for i := 1; i < len; i++ {
		if is_valid_string_valid_sentence(str[0:i], m) &&
			is_valid_string_valid_sentence(str[i:len], m) {
			fmt.Println(str[0:i], str[i:len])
			return true
		}
	}
	return false
}
func main() {
	path := "/etc/dictionaries-common/words"
	m := ReadDict(path)
	//fmt.Println(m)
	fmt.Println(is_valid_string_valid_sentence("fdsfdcatdogcat", m))
}
