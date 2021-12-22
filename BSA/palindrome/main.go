package main

import "fmt"

func palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
func main() {
	var kata string
	fmt.Scan(&kata)
	result := palindrome(kata)
	fmt.Println(result)
}
