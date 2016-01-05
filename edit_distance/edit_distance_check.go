package main

import "fmt"
import "math"

func is_edit_distance_one(str1 string, str2 string, gone bool) bool {
	//if str1 == "" && str2 == "" {
	//	return true
	//} else if str1 == "" && len(str2) == 1 && gone {
	//	return false
	//} else if str1 == "" && len(str2) == 1 && (!gone) {
	//	return true
	//} else if str2 == "" && len(str1) == 1 && gone {
	//	return false
	//} else if str2 == "" && len(str2) == 1 && (!gone) {
	//	return true
	//} else if str1[0] == str2[0] {
	//	return is_edit_distance_one(str1[1:], str2[1:], gone)
	//} else {
	//	if gone {
	//		return false
	//	} else {
	//		return is_edit_distance_one(str1[1:], str2, true) || is_edit_distance_one(str1, str2[1:], true)
	//	}
	//}
	if len(str1) == len(str2) {
		diff := false
		for i := 0; i < len(str1); i++ {
			if str1[i] != str2[i] {
				if !diff {
					diff = true
				} else {
					return false
				}
			}
		}
		return true
	} else if math.Abs(float64(len(str1)-len(str2))) == 1 {
		var short, long string
		done := false
		if len(str1) > len(str2) {
			short, long = str2, str1
		} else {
			short, long = str1, str2
		}
		for i, j := 0, 0; i < len(short) && j < len(long); {
			if short[i] == long[j] {
				i++
				j++
			} else {
				if !done {
					j++
					done = true
				} else {
					return false
				}
			}
		}
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(is_edit_distance_one("ac", "abc", false))
}
